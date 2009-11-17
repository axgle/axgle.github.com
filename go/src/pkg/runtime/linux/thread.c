<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/linux/thread.c</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/linux/thread.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;
#include &#34;defs.h&#34;
#include &#34;signals.h&#34;
#include &#34;os.h&#34;

// Linux futex.
//
//	futexsleep(uint32 *addr, uint32 val)
//	futexwakeup(uint32 *addr)
//
// Futexsleep atomically checks if *addr == val and if so, sleeps on addr.
// Futexwakeup wakes up one thread sleeping on addr.
// Futexsleep is allowed to wake up spuriously.

enum
{
	FUTEX_WAIT = 0,
	FUTEX_WAKE = 1,

	EINTR = 4,
	EAGAIN = 11,
};

// TODO(rsc): I tried using 1&lt;&lt;40 here but futex woke up (-ETIMEDOUT).
// I wonder if the timespec that gets to the kernel
// actually has two 32-bit numbers in it, so that
// a 64-bit 1&lt;&lt;40 ends up being 0 seconds,
// 1&lt;&lt;8 nanoseconds.
static Timespec longtime =
{
	1&lt;&lt;30,	// 34 years
	0
};

// Atomically,
//	if(*addr == val) sleep
// Might be woken up spuriously; that&#39;s allowed.
static void
futexsleep(uint32 *addr, uint32 val)
{
	int32 ret;

	ret = futex(addr, FUTEX_WAIT, val, &amp;longtime, nil, 0);
	if(ret &gt;= 0 || ret == -EAGAIN || ret == -EINTR)
		return;

	prints(&#34;futexsleep addr=&#34;);
	runtime·printpointer(addr);
	prints(&#34; val=&#34;);
	runtime·printint(val);
	prints(&#34; returned &#34;);
	runtime·printint(ret);
	prints(&#34;\n&#34;);
	*(int32*)0x1005 = 0x1005;
}

// If any procs are sleeping on addr, wake up at least one.
static void
futexwakeup(uint32 *addr)
{
	int64 ret;

	ret = futex(addr, FUTEX_WAKE, 1, nil, nil, 0);

	if(ret &gt;= 0)
		return;

	// I don&#39;t know that futex wakeup can return
	// EAGAIN or EINTR, but if it does, it would be
	// safe to loop and call futex again.

	prints(&#34;futexwakeup addr=&#34;);
	runtime·printpointer(addr);
	prints(&#34; returned &#34;);
	runtime·printint(ret);
	prints(&#34;\n&#34;);
	*(int32*)0x1006 = 0x1006;
}


// Lock and unlock.
//
// The lock state is a single 32-bit word that holds
// a 31-bit count of threads waiting for the lock
// and a single bit (the low bit) saying whether the lock is held.
// The uncontended case runs entirely in user space.
// When contention is detected, we defer to the kernel (futex).
//
// A reminder: compare-and-swap cas(addr, old, new) does
//	if(*addr == old) { *addr = new; return 1; }
//	else return 0;
// but atomically.

static void
futexlock(Lock *l)
{
	uint32 v;

again:
	v = l-&gt;key;
	if((v&amp;1) == 0){
		if(cas(&amp;l-&gt;key, v, v|1)){
			// Lock wasn&#39;t held; we grabbed it.
			return;
		}
		goto again;
	}

	// Lock was held; try to add ourselves to the waiter count.
	if(!cas(&amp;l-&gt;key, v, v+2))
		goto again;

	// We&#39;re accounted for, now sleep in the kernel.
	//
	// We avoid the obvious lock/unlock race because
	// the kernel won&#39;t put us to sleep if l-&gt;key has
	// changed underfoot and is no longer v+2.
	//
	// We only really care that (v&amp;1) == 1 (the lock is held),
	// and in fact there is a futex variant that could
	// accomodate that check, but let&#39;s not get carried away.)
	futexsleep(&amp;l-&gt;key, v+2);

	// We&#39;re awake: remove ourselves from the count.
	for(;;){
		v = l-&gt;key;
		if(v &lt; 2)
			throw(&#34;bad lock key&#34;);
		if(cas(&amp;l-&gt;key, v, v-2))
			break;
	}

	// Try for the lock again.
	goto again;
}

static void
futexunlock(Lock *l)
{
	uint32 v;

	// Atomically get value and clear lock bit.
again:
	v = l-&gt;key;
	if((v&amp;1) == 0)
		throw(&#34;unlock of unlocked lock&#34;);
	if(!cas(&amp;l-&gt;key, v, v&amp;~1))
		goto again;

	// If there were waiters, wake one.
	if(v &amp; ~1)
		futexwakeup(&amp;l-&gt;key);
}

void
lock(Lock *l)
{
	if(m-&gt;locks &lt; 0)
		throw(&#34;lock count&#34;);
	m-&gt;locks++;
	futexlock(l);
}

void
unlock(Lock *l)
{
	m-&gt;locks--;
	if(m-&gt;locks &lt; 0)
		throw(&#34;lock count&#34;);
	futexunlock(l);
}


// One-time notifications.
//
// Since the lock/unlock implementation already
// takes care of sleeping in the kernel, we just reuse it.
// (But it&#39;s a weird use, so it gets its own interface.)
//
// We use a lock to represent the event:
// unlocked == event has happened.
// Thus the lock starts out locked, and to wait for the
// event you try to lock the lock.  To signal the event,
// you unlock the lock.

void
noteclear(Note *n)
{
	n-&gt;lock.key = 0;	// memset(n, 0, sizeof *n)
	futexlock(&amp;n-&gt;lock);
}

void
notewakeup(Note *n)
{
	futexunlock(&amp;n-&gt;lock);
}

void
notesleep(Note *n)
{
	futexlock(&amp;n-&gt;lock);
	futexunlock(&amp;n-&gt;lock);	// Let other sleepers find out too.
}


// Clone, the Linux rfork.
enum
{
	CLONE_VM = 0x100,
	CLONE_FS = 0x200,
	CLONE_FILES = 0x400,
	CLONE_SIGHAND = 0x800,
	CLONE_PTRACE = 0x2000,
	CLONE_VFORK = 0x4000,
	CLONE_PARENT = 0x8000,
	CLONE_THREAD = 0x10000,
	CLONE_NEWNS = 0x20000,
	CLONE_SYSVSEM = 0x40000,
	CLONE_SETTLS = 0x80000,
	CLONE_PARENT_SETTID = 0x100000,
	CLONE_CHILD_CLEARTID = 0x200000,
	CLONE_UNTRACED = 0x800000,
	CLONE_CHILD_SETTID = 0x1000000,
	CLONE_STOPPED = 0x2000000,
	CLONE_NEWUTS = 0x4000000,
	CLONE_NEWIPC = 0x8000000,
};

void
newosproc(M *m, G *g, void *stk, void (*fn)(void))
{
	int32 ret;
	int32 flags;

	/*
	 * note: strace gets confused if we use CLONE_PTRACE here.
	 */
	flags = CLONE_PARENT	/* getppid doesn&#39;t change in child */
		| CLONE_VM	/* share memory */
		| CLONE_FS	/* share cwd, etc */
		| CLONE_FILES	/* share fd table */
		| CLONE_SIGHAND	/* share sig handler table */
		| CLONE_THREAD	/* revisit - okay for now */
		;

	m-&gt;tls[0] = m-&gt;id;	// so 386 asm can find it
	if(0){
		printf(&#34;newosproc stk=%p m=%p g=%p fn=%p clone=%p id=%d/%d ostk=%p\n&#34;,
			stk, m, g, fn, clone, m-&gt;id, m-&gt;tls[0], &amp;m);
	}

	ret = clone(flags, stk, m, g, fn);

	if(ret &lt; 0)
		*(int32*)123 = 123;
}

void
osinit(void)
{
}

// Called to initialize a new m (including the bootstrap m).
void
minit(void)
{
	// Initialize signal handling.
	m-&gt;gsignal = malg(32*1024);	// OS X wants &gt;=8K, Linux &gt;=2K
	signalstack(m-&gt;gsignal-&gt;stackguard, 32*1024);
}
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->

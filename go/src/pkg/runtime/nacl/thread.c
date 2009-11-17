<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/nacl/thread.c</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/nacl/thread.c</h1>

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
#include &#34;os.h&#34;

int8 *goos = &#34;nacl&#34;;

// Thread-safe allocation of a mutex.
// (The name sema is left over from the Darwin implementation.
// Native Client implements semaphores too, but it is just a shim
// over the host implementation, which on some hosts imposes a very
// low limit on how many semaphores can be created.)
//
// Psema points at a mutex descriptor.
// It starts out zero, meaning no mutex.
// Fill it in, being careful of others calling initsema
// simultaneously.
static void
initsema(uint32 *psema)
{
	uint32 sema;

	if(*psema != 0)	// already have one
		return;

	sema = mutex_create();
	if((int32)sema &lt; 0) {
		printf(&#34;mutex_create failed\n&#34;);
		breakpoint();
	}
	// mutex_create returns a file descriptor;
	// shift it up and add the 1 bit so that can
	// distinguish unintialized from fd 0.
	sema = (sema&lt;&lt;1) | 1;
	if(!cas(psema, 0, sema)){
		// Someone else filled it in.  Use theirs.
		close(sema);
		return;
	}
}

// Lock and unlock.
// Defer entirely to Native Client.
// The expense of a call into Native Client is more like
// a function call than a system call, so as long as the
// Native Client lock implementation is good, we can&#39;t
// do better ourselves.

static void
xlock(int32 fd)
{
	if(mutex_lock(fd) &lt; 0) {
		printf(&#34;mutex_lock failed\n&#34;);
		breakpoint();
	}
}

static void
xunlock(int32 fd)
{
	if(mutex_unlock(fd) &lt; 0) {
		printf(&#34;mutex_lock failed\n&#34;);
		breakpoint();
	}
}

void
lock(Lock *l)
{
	if(m-&gt;locks &lt; 0)
		throw(&#34;lock count&#34;);
	m-&gt;locks++;
	if(l-&gt;sema == 0)
		initsema(&amp;l-&gt;sema);
	xlock(l-&gt;sema&gt;&gt;1);
}

void
unlock(Lock *l)
{
	m-&gt;locks--;
	if(m-&gt;locks &lt; 0)
		throw(&#34;lock count&#34;);
	xunlock(l-&gt;sema&gt;&gt;1);
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
//
// Native Client does not require that the thread acquiring
// a lock be the thread that releases the lock, so this is safe.

void
noteclear(Note *n)
{
	if(n-&gt;lock.sema == 0)
		initsema(&amp;n-&gt;lock.sema);
	xlock(n-&gt;lock.sema&gt;&gt;1);
}

void
notewakeup(Note *n)
{
	if(n-&gt;lock.sema == 0) {
		printf(&#34;notewakeup without noteclear&#34;);
		breakpoint();
	}
	xunlock(n-&gt;lock.sema&gt;&gt;1);
}

void
notesleep(Note *n)
{
	if(n-&gt;lock.sema == 0) {
		printf(&#34;notesleep without noteclear&#34;);
		breakpoint();
	}
	xlock(n-&gt;lock.sema&gt;&gt;1);
	xunlock(n-&gt;lock.sema&gt;&gt;1);	// Let other sleepers find out too.
}

void
newosproc(M *m, G *g, void *stk, void (*fn)(void))
{
	void **vstk;

	// I wish every OS made thread creation this easy.
	m-&gt;tls[0] = (uint32)g;
	m-&gt;tls[1] = (uint32)m;
	vstk = stk;
	*--vstk = nil;
	if(thread_create(fn, vstk, m-&gt;tls, sizeof m-&gt;tls) &lt; 0) {
		printf(&#34;thread_create failed\n&#34;);
		breakpoint();
	}
}

void
osinit(void)
{
}

// Called to initialize a new m (including the bootstrap m).
void
minit(void)
{
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

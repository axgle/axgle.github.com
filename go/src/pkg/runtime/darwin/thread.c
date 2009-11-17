<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/darwin/thread.c</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/darwin/thread.c</h1>

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

static void
unimplemented(int8 *name)
{
	prints(name);
	prints(&#34; not implemented\n&#34;);
	*(int32*)1231 = 1231;
}

// Thread-safe allocation of a semaphore.
// Psema points at a kernel semaphore key.
// It starts out zero, meaning no semaphore.
// Fill it in, being careful of others calling initsema
// simultaneously.
static void
initsema(uint32 *psema)
{
	uint32 sema;

	if(*psema != 0)	// already have one
		return;

	sema = mach_semcreate();
	if(!cas(psema, 0, sema)){
		// Someone else filled it in.  Use theirs.
		mach_semdestroy(sema);
		return;
	}
}


// Atomic add and return new value.
static uint32
xadd(uint32 volatile *val, int32 delta)
{
	uint32 oval, nval;

	for(;;){
		oval = *val;
		nval = oval + delta;
		if(cas(val, oval, nval))
			return nval;
	}
}


// Blocking locks.

// Implement Locks, using semaphores.
// l-&gt;key is the number of threads who want the lock.
// In a race, one thread increments l-&gt;key from 0 to 1
// and the others increment it from &gt;0 to &gt;1.  The thread
// who does the 0-&gt;1 increment gets the lock, and the
// others wait on the semaphore.  When the 0-&gt;1 thread
// releases the lock by decrementing l-&gt;key, l-&gt;key will
// be &gt;0, so it will increment the semaphore to wake up
// one of the others.  This is the same algorithm used
// in Plan 9&#39;s user-level locks.
//
// Note that semaphores are never destroyed (the kernel
// will clean up when the process exits).  We assume for now
// that Locks are only used for long-lived structures like M and G.

void
lock(Lock *l)
{
	if(m-&gt;locks &lt; 0)
		throw(&#34;lock count&#34;);
	m-&gt;locks++;

	// Allocate semaphore if needed.
	if(l-&gt;sema == 0)
		initsema(&amp;l-&gt;sema);

	if(xadd(&amp;l-&gt;key, 1) &gt; 1)	// someone else has it; wait
		mach_semacquire(l-&gt;sema);
}

void
unlock(Lock *l)
{
	m-&gt;locks--;
	if(m-&gt;locks &lt; 0)
		throw(&#34;lock count&#34;);

	if(xadd(&amp;l-&gt;key, -1) &gt; 0)	// someone else is waiting
		mach_semrelease(l-&gt;sema);
}


// User-level semaphore implementation:
// try to do the operations in user space on u,
// but when it&#39;s time to block, fall back on the kernel semaphore k.
// This is the same algorithm used in Plan 9.
void
usemacquire(Usema *s)
{
	if((int32)xadd(&amp;s-&gt;u, -1) &lt; 0)
		mach_semacquire(s-&gt;k);
}

void
usemrelease(Usema *s)
{
	if((int32)xadd(&amp;s-&gt;u, 1) &lt;= 0)
		mach_semrelease(s-&gt;k);
}


// Event notifications.
void
noteclear(Note *n)
{
	n-&gt;wakeup = 0;
}

void
notesleep(Note *n)
{
	if(n-&gt;sema.k == 0)
		initsema(&amp;n-&gt;sema.k);
	while(!n-&gt;wakeup)
		usemacquire(&amp;n-&gt;sema);
}

void
notewakeup(Note *n)
{
	if(n-&gt;sema.k == 0)
		initsema(&amp;n-&gt;sema.k);
	n-&gt;wakeup = 1;
	usemrelease(&amp;n-&gt;sema);
}


// BSD interface for threading.
void
osinit(void)
{
	// Register our thread-creation callback (see {amd64,386}/sys.s)
	// but only if we&#39;re not using cgo.  If we are using cgo we need
	// to let the C pthread libary install its own thread-creation callback.
	extern void (*libcgo_thread_start)(void*);
	if(libcgo_thread_start == nil)
		bsdthread_register();
}

void
newosproc(M *m, G *g, void *stk, void (*fn)(void))
{
	m-&gt;tls[0] = m-&gt;id;	// so 386 asm can find it
	if(0){
		printf(&#34;newosproc stk=%p m=%p g=%p fn=%p id=%d/%d ostk=%p\n&#34;,
			stk, m, g, fn, m-&gt;id, m-&gt;tls[0], &amp;m);
	}
	bsdthread_create(stk, m, g, fn);
}

// Called to initialize a new m (including the bootstrap m).
void
minit(void)
{
	// Initialize signal handling.
	m-&gt;gsignal = malg(32*1024);	// OS X wants &gt;=8K, Linux &gt;=2K
	signalstack(m-&gt;gsignal-&gt;stackguard, 32*1024);
}

// Mach IPC, to get at semaphores
// Definitions are in /usr/include/mach on a Mac.

static void
macherror(int32 r, int8 *fn)
{
	printf(&#34;mach error %s: %d\n&#34;, fn, r);
	throw(&#34;mach error&#34;);
}

enum
{
	DebugMach = 0
};

static MachNDR zerondr;

#define MACH_MSGH_BITS(a, b) ((a) | ((b)&lt;&lt;8))

static int32
mach_msg(MachHeader *h,
	int32 op,
	uint32 send_size,
	uint32 rcv_size,
	uint32 rcv_name,
	uint32 timeout,
	uint32 notify)
{
	// TODO: Loop on interrupt.
	return mach_msg_trap(h, op, send_size, rcv_size, rcv_name, timeout, notify);
}

// Mach RPC (MIG)

enum
{
	MinMachMsg = 48,
	Reply = 100,
};

#pragma pack on
typedef struct CodeMsg CodeMsg;
struct CodeMsg
{
	MachHeader h;
	MachNDR NDR;
	int32 code;
};
#pragma pack off

static int32
machcall(MachHeader *h, int32 maxsize, int32 rxsize)
{
	uint32 *p;
	int32 i, ret, id;
	uint32 port;
	CodeMsg *c;

	if((port = m-&gt;machport) == 0){
		port = mach_reply_port();
		m-&gt;machport = port;
	}

	h-&gt;msgh_bits |= MACH_MSGH_BITS(MACH_MSG_TYPE_COPY_SEND, MACH_MSG_TYPE_MAKE_SEND_ONCE);
	h-&gt;msgh_local_port = port;
	h-&gt;msgh_reserved = 0;
	id = h-&gt;msgh_id;

	if(DebugMach){
		p = (uint32*)h;
		prints(&#34;send:\t&#34;);
		for(i=0; i&lt;h-&gt;msgh_size/sizeof(p[0]); i++){
			prints(&#34; &#34;);
			runtime·printpointer((void*)p[i]);
			if(i%8 == 7)
				prints(&#34;\n\t&#34;);
		}
		if(i%8)
			prints(&#34;\n&#34;);
	}

	ret = mach_msg(h, MACH_SEND_MSG|MACH_RCV_MSG,
		h-&gt;msgh_size, maxsize, port, 0, 0);
	if(ret != 0){
		if(DebugMach){
			prints(&#34;mach_msg error &#34;);
			runtime·printint(ret);
			prints(&#34;\n&#34;);
		}
		return ret;
	}

	if(DebugMach){
		p = (uint32*)h;
		prints(&#34;recv:\t&#34;);
		for(i=0; i&lt;h-&gt;msgh_size/sizeof(p[0]); i++){
			prints(&#34; &#34;);
			runtime·printpointer((void*)p[i]);
			if(i%8 == 7)
				prints(&#34;\n\t&#34;);
		}
		if(i%8)
			prints(&#34;\n&#34;);
	}

	if(h-&gt;msgh_id != id+Reply){
		if(DebugMach){
			prints(&#34;mach_msg reply id mismatch &#34;);
			runtime·printint(h-&gt;msgh_id);
			prints(&#34; != &#34;);
			runtime·printint(id+Reply);
			prints(&#34;\n&#34;);
		}
		return -303;	// MIG_REPLY_MISMATCH
	}

	// Look for a response giving the return value.
	// Any call can send this back with an error,
	// and some calls only have return values so they
	// send it back on success too.  I don&#39;t quite see how
	// you know it&#39;s one of these and not the full response
	// format, so just look if the message is right.
	c = (CodeMsg*)h;
	if(h-&gt;msgh_size == sizeof(CodeMsg)
	&amp;&amp; !(h-&gt;msgh_bits &amp; MACH_MSGH_BITS_COMPLEX)){
		if(DebugMach){
			prints(&#34;mig result &#34;);
			runtime·printint(c-&gt;code);
			prints(&#34;\n&#34;);
		}
		return c-&gt;code;
	}

	if(h-&gt;msgh_size != rxsize){
		if(DebugMach){
			prints(&#34;mach_msg reply size mismatch &#34;);
			runtime·printint(h-&gt;msgh_size);
			prints(&#34; != &#34;);
			runtime·printint(rxsize);
			prints(&#34;\n&#34;);
		}
		return -307;	// MIG_ARRAY_TOO_LARGE
	}

	return 0;
}


// Semaphores!

enum
{
	Tmach_semcreate = 3418,
	Rmach_semcreate = Tmach_semcreate + Reply,

	Tmach_semdestroy = 3419,
	Rmach_semdestroy = Tmach_semdestroy + Reply,

	// Mach calls that get interrupted by Unix signals
	// return this error code.  We retry them.
	KERN_ABORTED = 14,
};

typedef struct Tmach_semcreateMsg Tmach_semcreateMsg;
typedef struct Rmach_semcreateMsg Rmach_semcreateMsg;
typedef struct Tmach_semdestroyMsg Tmach_semdestroyMsg;
// Rmach_semdestroyMsg = CodeMsg

#pragma pack on
struct Tmach_semcreateMsg
{
	MachHeader h;
	MachNDR ndr;
	int32 policy;
	int32 value;
};

struct Rmach_semcreateMsg
{
	MachHeader h;
	MachBody body;
	MachPort semaphore;
};

struct Tmach_semdestroyMsg
{
	MachHeader h;
	MachBody body;
	MachPort semaphore;
};
#pragma pack off

uint32
mach_semcreate(void)
{
	union {
		Tmach_semcreateMsg tx;
		Rmach_semcreateMsg rx;
		uint8 pad[MinMachMsg];
	} m;
	int32 r;

	m.tx.h.msgh_bits = 0;
	m.tx.h.msgh_size = sizeof(m.tx);
	m.tx.h.msgh_remote_port = mach_task_self();
	m.tx.h.msgh_id = Tmach_semcreate;
	m.tx.ndr = zerondr;

	m.tx.policy = 0;	// 0 = SYNC_POLICY_FIFO
	m.tx.value = 0;

	while((r = machcall(&amp;m.tx.h, sizeof m, sizeof(m.rx))) != 0){
		if(r == KERN_ABORTED)	// interrupted
			continue;
		macherror(r, &#34;semaphore_create&#34;);
	}
	if(m.rx.body.msgh_descriptor_count != 1)
		unimplemented(&#34;mach_semcreate desc count&#34;);
	return m.rx.semaphore.name;
}

void
mach_semdestroy(uint32 sem)
{
	union {
		Tmach_semdestroyMsg tx;
		uint8 pad[MinMachMsg];
	} m;
	int32 r;

	m.tx.h.msgh_bits = MACH_MSGH_BITS_COMPLEX;
	m.tx.h.msgh_size = sizeof(m.tx);
	m.tx.h.msgh_remote_port = mach_task_self();
	m.tx.h.msgh_id = Tmach_semdestroy;
	m.tx.body.msgh_descriptor_count = 1;
	m.tx.semaphore.name = sem;
	m.tx.semaphore.disposition = MACH_MSG_TYPE_MOVE_SEND;
	m.tx.semaphore.type = 0;

	while((r = machcall(&amp;m.tx.h, sizeof m, 0)) != 0){
		if(r == KERN_ABORTED)	// interrupted
			continue;
		macherror(r, &#34;semaphore_destroy&#34;);
	}
}

// The other calls have simple system call traps in sys.s
int32 mach_semaphore_wait(uint32 sema);
int32 mach_semaphore_timedwait(uint32 sema, uint32 sec, uint32 nsec);
int32 mach_semaphore_signal(uint32 sema);
int32 mach_semaphore_signal_all(uint32 sema);

void
mach_semacquire(uint32 sem)
{
	int32 r;

	while((r = mach_semaphore_wait(sem)) != 0) {
		if(r == KERN_ABORTED)	// interrupted
			continue;
		macherror(r, &#34;semaphore_wait&#34;);
	}
}

void
mach_semrelease(uint32 sem)
{
	int32 r;

	while((r = mach_semaphore_signal(sem)) != 0) {
		if(r == KERN_ABORTED)	// interrupted
			continue;
		macherror(r, &#34;semaphore_signal&#34;);
	}
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

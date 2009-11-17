<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/darwin.c</title>

  <link rel="stylesheet" type="text/css" href="../../doc/style.css">
  <script type="text/javascript" src="../../doc/godocs.js"></script>

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
        <a href="../../index.html"><img src="../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../doc/install.html">Install Go</a></li>
    <li><a href="../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/libmach/darwin.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
//	Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the &#34;Software&#34;), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#define __DARWIN_UNIX03 0

#include &lt;u.h&gt;
#include &lt;sys/ptrace.h&gt;
#include &lt;sys/signal.h&gt;
#include &lt;mach/mach.h&gt;
#include &lt;mach/mach_traps.h&gt;
#include &lt;errno.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;mach.h&gt;
#define Ureg Ureg32
#include &lt;ureg_x86.h&gt;
#undef Ureg
#define Ureg Ureg64
#include &lt;ureg_amd64.h&gt;
#undef Ureg
#undef waitpid	/* want Unix waitpid, not Plan 9 */

typedef struct Ureg32 Ureg32;
typedef struct Ureg64 Ureg64;

extern mach_port_t mach_reply_port(void);	// should be in system headers, is not

// Mach-error wrapper.
// Takes a mach return code and converts it into 0 / -1,
// setting errstr when it returns -1.

static struct {
	int code;
	char *name;
} macherr[] = {
	KERN_INVALID_ADDRESS,	&#34;invalid address&#34;,
	KERN_PROTECTION_FAILURE,	&#34;protection failure&#34;,
	KERN_NO_SPACE,	&#34;no space&#34;,
	KERN_INVALID_ARGUMENT,	&#34;invalid argument&#34;,
	KERN_FAILURE,	&#34;failure&#34;,
	KERN_RESOURCE_SHORTAGE,	&#34;resource shortage&#34;,
	KERN_NOT_RECEIVER,	&#34;not receiver&#34;,
	KERN_NO_ACCESS,	&#34;no access&#34;,
	KERN_MEMORY_FAILURE,	&#34;memory failure&#34;,
	KERN_MEMORY_ERROR,	&#34;memory error&#34;,
	KERN_ALREADY_IN_SET,	&#34;already in set&#34;,
	KERN_NOT_IN_SET,	&#34;not in set&#34;,
	KERN_NAME_EXISTS,	&#34;name exists&#34;,
	KERN_ABORTED,	&#34;aborted&#34;,
	KERN_INVALID_NAME,	&#34;invalid name&#34;,
	KERN_INVALID_TASK,	&#34;invalid task&#34;,
	KERN_INVALID_RIGHT,	&#34;invalid right&#34;,
	KERN_INVALID_VALUE,	&#34;invalid value&#34;,
	KERN_UREFS_OVERFLOW,	&#34;urefs overflow&#34;,
	KERN_INVALID_CAPABILITY,	&#34;invalid capability&#34;,
	KERN_RIGHT_EXISTS,	&#34;right exists&#34;,
	KERN_INVALID_HOST,	&#34;invalid host&#34;,
	KERN_MEMORY_PRESENT,	&#34;memory present&#34;,
	KERN_MEMORY_DATA_MOVED,	&#34;memory data moved&#34;,
	KERN_MEMORY_RESTART_COPY,	&#34;memory restart copy&#34;,
	KERN_INVALID_PROCESSOR_SET,	&#34;invalid processor set&#34;,
	KERN_POLICY_LIMIT,	&#34;policy limit&#34;,
	KERN_INVALID_POLICY,	&#34;invalid policy&#34;,
	KERN_INVALID_OBJECT,	&#34;invalid object&#34;,
	KERN_ALREADY_WAITING,	&#34;already waiting&#34;,
	KERN_DEFAULT_SET,	&#34;default set&#34;,
	KERN_EXCEPTION_PROTECTED,	&#34;exception protected&#34;,
	KERN_INVALID_LEDGER,	&#34;invalid ledger&#34;,
	KERN_INVALID_MEMORY_CONTROL,	&#34;invalid memory control&#34;,
	KERN_INVALID_SECURITY,	&#34;invalid security&#34;,
	KERN_NOT_DEPRESSED,	&#34;not depressed&#34;,
	KERN_TERMINATED,	&#34;terminated&#34;,
	KERN_LOCK_SET_DESTROYED,	&#34;lock set destroyed&#34;,
	KERN_LOCK_UNSTABLE,	&#34;lock unstable&#34;,
	KERN_LOCK_OWNED,	&#34;lock owned&#34;,
	KERN_LOCK_OWNED_SELF,	&#34;lock owned self&#34;,
	KERN_SEMAPHORE_DESTROYED,	&#34;semaphore destroyed&#34;,
	KERN_RPC_SERVER_TERMINATED,	&#34;rpc server terminated&#34;,
	KERN_RPC_TERMINATE_ORPHAN,	&#34;rpc terminate orphan&#34;,
	KERN_RPC_CONTINUE_ORPHAN,	&#34;rpc continue orphan&#34;,
	KERN_NOT_SUPPORTED,	&#34;not supported&#34;,
	KERN_NODE_DOWN,	&#34;node down&#34;,
	KERN_NOT_WAITING,	&#34;not waiting&#34;,
	KERN_OPERATION_TIMED_OUT,	&#34;operation timed out&#34;,
	KERN_RETURN_MAX,	&#34;return max&#34;,

	MACH_SEND_IN_PROGRESS,	&#34;send in progress&#34;,
	MACH_SEND_INVALID_DATA,	&#34;send invalid data&#34;,
	MACH_SEND_INVALID_DEST,	&#34;send invalid dest&#34;,
	MACH_SEND_TIMED_OUT,	&#34;send timed out&#34;,
	MACH_SEND_INTERRUPTED,	&#34;send interrupted&#34;,
	MACH_SEND_MSG_TOO_SMALL,	&#34;send msg too small&#34;,
	MACH_SEND_INVALID_REPLY,	&#34;send invalid reply&#34;,
	MACH_SEND_INVALID_RIGHT,	&#34;send invalid right&#34;,
	MACH_SEND_INVALID_NOTIFY,	&#34;send invalid notify&#34;,
	MACH_SEND_INVALID_MEMORY,	&#34;send invalid memory&#34;,
	MACH_SEND_NO_BUFFER,	&#34;send no buffer&#34;,
	MACH_SEND_TOO_LARGE,	&#34;send too large&#34;,
	MACH_SEND_INVALID_TYPE,	&#34;send invalid type&#34;,
	MACH_SEND_INVALID_HEADER,	&#34;send invalid header&#34;,
	MACH_SEND_INVALID_TRAILER,	&#34;send invalid trailer&#34;,
	MACH_SEND_INVALID_RT_OOL_SIZE,	&#34;send invalid rt ool size&#34;,
	MACH_RCV_IN_PROGRESS,	&#34;rcv in progress&#34;,
	MACH_RCV_INVALID_NAME,	&#34;rcv invalid name&#34;,
	MACH_RCV_TIMED_OUT,	&#34;rcv timed out&#34;,
	MACH_RCV_TOO_LARGE,	&#34;rcv too large&#34;,
	MACH_RCV_INTERRUPTED,	&#34;rcv interrupted&#34;,
	MACH_RCV_PORT_CHANGED,	&#34;rcv port changed&#34;,
	MACH_RCV_INVALID_NOTIFY,	&#34;rcv invalid notify&#34;,
	MACH_RCV_INVALID_DATA,	&#34;rcv invalid data&#34;,
	MACH_RCV_PORT_DIED,	&#34;rcv port died&#34;,
	MACH_RCV_IN_SET,	&#34;rcv in set&#34;,
	MACH_RCV_HEADER_ERROR,	&#34;rcv header error&#34;,
	MACH_RCV_BODY_ERROR,	&#34;rcv body error&#34;,
	MACH_RCV_INVALID_TYPE,	&#34;rcv invalid type&#34;,
	MACH_RCV_SCATTER_SMALL,	&#34;rcv scatter small&#34;,
	MACH_RCV_INVALID_TRAILER,	&#34;rcv invalid trailer&#34;,
	MACH_RCV_IN_PROGRESS_TIMED,	&#34;rcv in progress timed&#34;,

	MIG_TYPE_ERROR,	&#34;mig type error&#34;,
	MIG_REPLY_MISMATCH,	&#34;mig reply mismatch&#34;,
	MIG_REMOTE_ERROR,	&#34;mig remote error&#34;,
	MIG_BAD_ID,	&#34;mig bad id&#34;,
	MIG_BAD_ARGUMENTS,	&#34;mig bad arguments&#34;,
	MIG_NO_REPLY,	&#34;mig no reply&#34;,
	MIG_EXCEPTION,	&#34;mig exception&#34;,
	MIG_ARRAY_TOO_LARGE,	&#34;mig array too large&#34;,
	MIG_SERVER_DIED,	&#34;server died&#34;,
	MIG_TRAILER_ERROR,	&#34;trailer has an unknown format&#34;,
};

static int
me(kern_return_t r)
{
	int i;

	if(r == 0)
		return 0;

	for(i=0; i&lt;nelem(macherr); i++){
		if(r == macherr[i].code){
			werrstr(&#34;%s&#34;, macherr[i].name);
			return -1;
		}
	}
	werrstr(&#34;mach error %#x&#34;, r);
	return -1;
}

// Plan 9 and Linux do not distinguish between
// process ids and thread ids, so the interface here doesn&#39;t either.
// Unfortunately, Mach has three kinds of identifiers: process ids,
// handles to tasks (processes), and handles to threads within a
// process.  All of them are small integers.
//
// To accomodate Mach, we employ a clumsy hack: in this interface,
// if you pass in a positive number, that&#39;s a process id.
// If you pass in a negative number, that identifies a thread that
// has been previously returned by procthreadpids (it indexes
// into the Thread table below).

// Table of threads we have handles for.
typedef struct Thread Thread;
struct Thread
{
	int pid;
	mach_port_t task;
	mach_port_t thread;
	int stopped;
	int exc;
	int code[10];
	Map *map;
};
static Thread thr[1000];
static int nthr;
static pthread_mutex_t mu;
static pthread_cond_t cond;
static void* excthread(void*);
static void* waitthread(void*);
static mach_port_t excport;

enum {
	ExcMask = EXC_MASK_BAD_ACCESS |
		EXC_MASK_BAD_INSTRUCTION |
		EXC_MASK_ARITHMETIC |
		EXC_MASK_BREAKPOINT |
		EXC_MASK_SOFTWARE
};

// Add process pid to the thread table.
// If it&#39;s already there, don&#39;t re-add it (unless force != 0).
static Thread*
addpid(int pid, int force)
{
	int i, j;
	mach_port_t task;
	mach_port_t *thread;
	uint nthread;
	Thread *ret;
	static int first = 1;

	if(first){
		// Allocate a port for exception messages and
		// send all thread exceptions to that port.
		// The excthread reads that port and signals
		// us if we are waiting on that thread.
		pthread_t p;

		excport = mach_reply_port();
		pthread_mutex_init(&amp;mu, nil);
		pthread_cond_init(&amp;cond, nil);
		pthread_create(&amp;p, nil, excthread, nil);
		pthread_create(&amp;p, nil, waitthread, (void*)(uintptr)pid);
		first = 0;
	}

	if(!force){
		for(i=0; i&lt;nthr; i++)
			if(thr[i].pid == pid)
				return &amp;thr[i];
	}
	if(me(task_for_pid(mach_task_self(), pid, &amp;task)) &lt; 0)
		return nil;
	if(me(task_threads(task, &amp;thread, &amp;nthread)) &lt; 0)
		return nil;
	mach_port_insert_right(mach_task_self(), excport, excport, MACH_MSG_TYPE_MAKE_SEND);
	if(me(task_set_exception_ports(task, ExcMask,
			excport, EXCEPTION_DEFAULT, MACHINE_THREAD_STATE)) &lt; 0){
		fprint(2, &#34;warning: cannot set excport: %r\n&#34;);
	}
	ret = nil;
	for(j=0; j&lt;nthread; j++){
		if(force){
			// If we&#39;re forcing a refresh, don&#39;t re-add existing threads.
			for(i=0; i&lt;nthr; i++)
				if(thr[i].pid == pid &amp;&amp; thr[i].thread == thread[j]){
					if(ret == nil)
						ret = &amp;thr[i];
					goto skip;
				}
		}
		if(nthr &gt;= nelem(thr))
			return nil;
		// TODO: We probably should save the old thread exception
		// ports for each bit and then put them back when we exit.
		// Probably the BSD signal handlers have put stuff there.
		mach_port_insert_right(mach_task_self(), excport, excport, MACH_MSG_TYPE_MAKE_SEND);
		if(me(thread_set_exception_ports(thread[j], ExcMask,
				excport, EXCEPTION_DEFAULT, MACHINE_THREAD_STATE)) &lt; 0){
			fprint(2, &#34;warning: cannot set excport: %r\n&#34;);
		}
		thr[nthr].pid = pid;
		thr[nthr].task = task;
		thr[nthr].thread = thread[j];
		if(ret == nil)
			ret = &amp;thr[nthr];
		nthr++;
	skip:;
	}
	return ret;
}

static Thread*
idtotable(int id)
{
	if(id &gt;= 0)
		return addpid(id, 1);

	id = -(id+1);
	if(id &gt;= nthr)
		return nil;
	return &amp;thr[id];
}

/*
static int
idtopid(int id)
{
	Thread *t;

	if((t = idtotable(id)) == nil)
		return -1;
	return t-&gt;pid;
}
*/

static mach_port_t
idtotask(int id)
{
	Thread *t;

	if((t = idtotable(id)) == nil)
		return -1;
	return t-&gt;task;
}

static mach_port_t
idtothread(int id)
{
	Thread *t;

	if((t = idtotable(id)) == nil)
		return -1;
	return t-&gt;thread;
}

static int machsegrw(Map *map, Seg *seg, uvlong addr, void *v, uint n, int isr);
static int machregrw(Map *map, Seg *seg, uvlong addr, void *v, uint n, int isr);

Map*
attachproc(int id, Fhdr *fp)
{
	Thread *t;
	Map *map;

	if((t = idtotable(id)) == nil)
		return nil;
	if(t-&gt;map)
		return t-&gt;map;
	map = newmap(0, 4);
	if(!map)
		return nil;
	map-&gt;pid = -((t-thr) + 1);
	if(mach-&gt;regsize)
		setmap(map, -1, 0, mach-&gt;regsize, 0, &#34;regs&#34;, machregrw);
	setmap(map, -1, fp-&gt;txtaddr, fp-&gt;txtaddr+fp-&gt;txtsz, fp-&gt;txtaddr, &#34;*text&#34;, machsegrw);
	setmap(map, -1, fp-&gt;dataddr, mach-&gt;utop, fp-&gt;dataddr, &#34;*data&#34;, machsegrw);
	t-&gt;map = map;
	return map;
}

// Return list of ids for threads in id.
int
procthreadpids(int id, int *out, int nout)
{
	Thread *t;
	int i, n, pid;

	t = idtotable(id);
	if(t == nil)
		return -1;
	pid = t-&gt;pid;
	addpid(pid, 1);	// force refresh of thread list
	n = 0;
	for(i=0; i&lt;nthr; i++) {
		if(thr[i].pid == pid) {
			if(n &lt; nout)
				out[n] = -(i+1);
			n++;
		}
	}
	return n;
}

// Detach from proc.
// TODO(rsc): Perhaps should unsuspend any threads and clean-up the table.
void
detachproc(Map *m)
{
	free(m);
}

// Should return array of pending signals (notes)
// but don&#39;t know how to do that on OS X.
int
procnotes(int pid, char ***pnotes)
{
	*pnotes = 0;
	return 0;
}

// There must be a way to do this.  Gdb can do it.
// But I don&#39;t see, in the Apple gdb sources, how.
char*
proctextfile(int pid)
{
	return nil;
}

// Read/write from a Mach data segment.
static int
machsegrw(Map *map, Seg *seg, uvlong addr, void *v, uint n, int isr)
{
	mach_port_t task;
	int r;

	task = idtotask(map-&gt;pid);
	if(task == -1)
		return -1;

	if(isr){
		vm_size_t nn;
		nn = n;
		if(me(vm_read_overwrite(task, addr, n, (uintptr)v, &amp;nn)) &lt; 0)
			return -1;
		return nn;
	}else{
		r = vm_write(task, addr, (uintptr)v, n);
		if(r == KERN_INVALID_ADDRESS){
			// Happens when writing to text segment.
			// Change protections.
			if(me(vm_protect(task, addr, n, 0, VM_PROT_WRITE|VM_PROT_READ|VM_PROT_EXECUTE)) &lt; 0){
				fprint(2, &#34;vm_protect: %s\n&#34;, r);
				return -1;
			}
			r = vm_write(task, addr, (uintptr)v, n);
		}
		if(r != 0){
			me(r);
			return -1;
		}
		return n;
	}
}

// Convert Ureg offset to x86_thread_state32_t offset.
static int
go2darwin32(uvlong addr)
{
	switch(addr){
	case offsetof(Ureg32, ax):
		return offsetof(x86_thread_state32_t, eax);
	case offsetof(Ureg32, bx):
		return offsetof(x86_thread_state32_t, ebx);
	case offsetof(Ureg32, cx):
		return offsetof(x86_thread_state32_t, ecx);
	case offsetof(Ureg32, dx):
		return offsetof(x86_thread_state32_t, edx);
	case offsetof(Ureg32, si):
		return offsetof(x86_thread_state32_t, esi);
	case offsetof(Ureg32, di):
		return offsetof(x86_thread_state32_t, edi);
	case offsetof(Ureg32, bp):
		return offsetof(x86_thread_state32_t, ebp);
	case offsetof(Ureg32, fs):
		return offsetof(x86_thread_state32_t, fs);
	case offsetof(Ureg32, gs):
		return offsetof(x86_thread_state32_t, gs);
	case offsetof(Ureg32, pc):
		return offsetof(x86_thread_state32_t, eip);
	case offsetof(Ureg32, cs):
		return offsetof(x86_thread_state32_t, cs);
	case offsetof(Ureg32, flags):
		return offsetof(x86_thread_state32_t, eflags);
	case offsetof(Ureg32, sp):
		return offsetof(x86_thread_state32_t, esp);
	}
	return -1;
}

// Convert Ureg offset to x86_thread_state64_t offset.
static int
go2darwin64(uvlong addr)
{
	switch(addr){
	case offsetof(Ureg64, ax):
		return offsetof(x86_thread_state64_t, rax);
	case offsetof(Ureg64, bx):
		return offsetof(x86_thread_state64_t, rbx);
	case offsetof(Ureg64, cx):
		return offsetof(x86_thread_state64_t, rcx);
	case offsetof(Ureg64, dx):
		return offsetof(x86_thread_state64_t, rdx);
	case offsetof(Ureg64, si):
		return offsetof(x86_thread_state64_t, rsi);
	case offsetof(Ureg64, di):
		return offsetof(x86_thread_state64_t, rdi);
	case offsetof(Ureg64, bp):
		return offsetof(x86_thread_state64_t, rbp);
	case offsetof(Ureg64, r8):
		return offsetof(x86_thread_state64_t, r8);
	case offsetof(Ureg64, r9):
		return offsetof(x86_thread_state64_t, r9);
	case offsetof(Ureg64, r10):
		return offsetof(x86_thread_state64_t, r10);
	case offsetof(Ureg64, r11):
		return offsetof(x86_thread_state64_t, r11);
	case offsetof(Ureg64, r12):
		return offsetof(x86_thread_state64_t, r12);
	case offsetof(Ureg64, r13):
		return offsetof(x86_thread_state64_t, r13);
	case offsetof(Ureg64, r14):
		return offsetof(x86_thread_state64_t, r14);
	case offsetof(Ureg64, r15):
		return offsetof(x86_thread_state64_t, r15);
	case offsetof(Ureg64, fs):
		return offsetof(x86_thread_state64_t, fs);
	case offsetof(Ureg64, gs):
		return offsetof(x86_thread_state64_t, gs);
	case offsetof(Ureg64, ip):
		return offsetof(x86_thread_state64_t, rip);
	case offsetof(Ureg64, cs):
		return offsetof(x86_thread_state64_t, cs);
	case offsetof(Ureg64, flags):
		return offsetof(x86_thread_state64_t, rflags);
	case offsetof(Ureg64, sp):
		return offsetof(x86_thread_state64_t, rsp);
	}
	return -1;
}

extern Mach mi386;

// Read/write from fake register segment.
static int
machregrw(Map *map, Seg *seg, uvlong addr, void *v, uint n, int isr)
{
	uint nn, count, state;
	mach_port_t thread;
	int reg;
	char buf[100];
	union {
		x86_thread_state64_t reg64;
		x86_thread_state32_t reg32;
		uchar p[1];
	} u;
	uchar *p;

	if(n &gt; 8){
		werrstr(&#34;asked for %d-byte register&#34;, n);
		return -1;
	}

	thread = idtothread(map-&gt;pid);
	if(thread == -1){
		werrstr(&#34;no such id&#34;);
		return -1;
	}

	if(mach == &amp;mi386) {
		count = x86_THREAD_STATE32_COUNT;
		state = x86_THREAD_STATE32;
		if((reg = go2darwin32(addr)) &lt; 0 || reg+n &gt; sizeof u){
			if(isr){
				memset(v, 0, n);
				return 0;
			}
			werrstr(&#34;register %llud not available&#34;, addr);
			return -1;
		}
	} else {
		count = x86_THREAD_STATE64_COUNT;
		state = x86_THREAD_STATE64;
		if((reg = go2darwin64(addr)) &lt; 0 || reg+n &gt; sizeof u){
			if(isr){
				memset(v, 0, n);
				return 0;
			}
			werrstr(&#34;register %llud not available&#34;, addr);
			return -1;
		}
	}

	if(!isr &amp;&amp; me(thread_suspend(thread)) &lt; 0){
		werrstr(&#34;thread suspend %#x: %r&#34;, thread);
		return -1;
	}
	nn = count;
	if(me(thread_get_state(thread, state, (void*)u.p, &amp;nn)) &lt; 0){
		if(!isr)
			thread_resume(thread);
		rerrstr(buf, sizeof buf);
		if(strcmp(buf, &#34;send invalid dest&#34;) == 0)
			werrstr(&#34;process exited&#34;);
		else
			werrstr(&#34;thread_get_state: %r&#34;);
		return -1;
	}

	p = u.p+reg;
	if(isr)
		memmove(v, p, n);
	else{
		memmove(p, v, n);
		nn = count;
		if(me(thread_set_state(thread, state, (void*)u.p, nn)) &lt; 0){
			thread_resume(thread);
			werrstr(&#34;thread_set_state: %r&#34;);
			return -1;
		}

		if(me(thread_resume(thread)) &lt; 0){
			werrstr(&#34;thread_resume: %r&#34;);
			return -1;
		}
	}
	return 0;
}

enum
{
	FLAGS_TF = 0x100		// x86 single-step processor flag
};

// Is thread t suspended?
static int
threadstopped(Thread *t)
{
	struct thread_basic_info info;
	uint size;

	size = sizeof info;
	if(me(thread_info(t-&gt;thread, THREAD_BASIC_INFO, (thread_info_t)&amp;info, &amp;size)) &lt;  0){
		fprint(2, &#34;threadstopped thread_info %#x: %r\n&#34;);
		return 1;
	}
	return info.suspend_count &gt; 0;
}

// If thread t is suspended, start it up again.
// If singlestep is set, only let it execute one instruction.
static int
threadstart(Thread *t, int singlestep)
{
	int i;
	uint n;
	struct thread_basic_info info;

	if(!threadstopped(t))
		return 0;

	// Set or clear the processor single-step flag, as appropriate.
	if(mach == &amp;mi386) {
		x86_thread_state32_t regs;
		n = x86_THREAD_STATE32_COUNT;
		if(me(thread_get_state(t-&gt;thread, x86_THREAD_STATE32,
				(thread_state_t)&amp;regs,
				&amp;n)) &lt; 0)
			return -1;
		if(singlestep)
			regs.eflags |= FLAGS_TF;
		else
			regs.eflags &amp;= ~FLAGS_TF;
		if(me(thread_set_state(t-&gt;thread, x86_THREAD_STATE32,
				(thread_state_t)&amp;regs,
				x86_THREAD_STATE32_COUNT)) &lt; 0)
			return -1;
	} else {
		x86_thread_state64_t regs;
		n = x86_THREAD_STATE64_COUNT;
		if(me(thread_get_state(t-&gt;thread, x86_THREAD_STATE64,
				(thread_state_t)&amp;regs,
				&amp;n)) &lt; 0)
			return -1;
		if(singlestep)
			regs.rflags |= FLAGS_TF;
		else
			regs.rflags &amp;= ~FLAGS_TF;
		if(me(thread_set_state(t-&gt;thread, x86_THREAD_STATE64,
				(thread_state_t)&amp;regs,
				x86_THREAD_STATE64_COUNT)) &lt; 0)
			return -1;
	}

	// Run.
	n = sizeof info;
	if(me(thread_info(t-&gt;thread, THREAD_BASIC_INFO, (thread_info_t)&amp;info, &amp;n)) &lt; 0)
		return -1;
	for(i=0; i&lt;info.suspend_count; i++)
		if(me(thread_resume(t-&gt;thread)) &lt; 0)
			return -1;
	return 0;
}

// Stop thread t.
static int
threadstop(Thread *t)
{
	if(threadstopped(t))
		return 0;
	if(me(thread_suspend(t-&gt;thread)) &lt; 0)
		return -1;
	return 0;
}

// Callback for exc_server below.  Called when a thread we are
// watching has an exception like hitting a breakpoint.
kern_return_t
catch_exception_raise(mach_port_t eport, mach_port_t thread,
	mach_port_t task, exception_type_t exception,
	exception_data_t code, mach_msg_type_number_t ncode)
{
	Thread *t;
	int i;

	t = nil;
	for(i=0; i&lt;nthr; i++){
		if(thr[i].thread == thread){
			t = &amp;thr[i];
			goto havet;
		}
	}
	if(nthr &gt; 0)
		addpid(thr[0].pid, 1);
	for(i=0; i&lt;nthr; i++){
		if(thr[i].thread == thread){
			t = &amp;thr[i];
			goto havet;
		}
	}
	fprint(2, &#34;did not find thread in catch_exception_raise\n&#34;);
	return KERN_SUCCESS;	// let thread continue

havet:
	t-&gt;exc = exception;
	if(ncode &gt; nelem(t-&gt;code))
		ncode = nelem(t-&gt;code);
	memmove(t-&gt;code, code, ncode*sizeof t-&gt;code[0]);

	// Suspend thread, so that we can look at it &amp; restart it later.
	if(me(thread_suspend(thread)) &lt; 0)
		fprint(2, &#34;catch_exception_raise thread_suspend: %r\n&#34;);

	// Synchronize with waitstop below.
	pthread_mutex_lock(&amp;mu);
	pthread_cond_broadcast(&amp;cond);
	pthread_mutex_unlock(&amp;mu);

	return KERN_SUCCESS;
}

// Exception watching thread, started in addpid above.
static void*
excthread(void *v)
{
	extern boolean_t exc_server();
	mach_msg_server(exc_server, 2048, excport, 0);
	return 0;
}

// Wait for pid to exit.
static int exited;
static void*
waitthread(void *v)
{
	int pid, status;

	pid = (int)(uintptr)v;
	waitpid(pid, &amp;status, 0);
	exited = 1;
	// Synchronize with waitstop below.
	pthread_mutex_lock(&amp;mu);
	pthread_cond_broadcast(&amp;cond);
	pthread_mutex_unlock(&amp;mu);
	return nil;
}

// Wait for thread t to stop.
static int
waitstop(Thread *t)
{
	pthread_mutex_lock(&amp;mu);
	while(!exited &amp;&amp; !threadstopped(t))
		pthread_cond_wait(&amp;cond, &amp;mu);
	pthread_mutex_unlock(&amp;mu);
	return 0;
}

int
ctlproc(int id, char *msg)
{
	Thread *t;
	int status;

	// Hang/attached dance is for debugging newly exec&#39;ed programs.
	// After fork, the child does ctlproc(&#34;hang&#34;) before exec,
	// and the parent does ctlproc(&#34;attached&#34;) and then waitstop.
	// Using these requires the BSD ptrace interface, unlike everything
	// else we do, which uses only the Mach interface.  Our goal here
	// is to do as little as possible using ptrace and then flip over to Mach.

	if(strcmp(msg, &#34;hang&#34;) == 0)
		return ptrace(PT_TRACE_ME, 0, 0, 0);

	if(strcmp(msg, &#34;attached&#34;) == 0){
		// The pid &#34;id&#34; has done a ctlproc &#34;hang&#34; and then
		// exec, so we should find it stoppped just before exec
		// of the new program.
		#undef waitpid
		if(waitpid(id, &amp;status, WUNTRACED) &lt; 0){
			fprint(2, &#34;ctlproc attached waitpid: %r\n&#34;);
			return -1;
		}
		if(WIFEXITED(status) || !WIFSTOPPED(status)){
			fprint(2, &#34;ctlproc attached: bad process state\n&#34;);
			return -1;
		}

		// Find Mach thread for pid and suspend it.
		t = addpid(id, 1);
		if(t == nil)
			return -1;
		if(me(thread_suspend(t-&gt;thread)) &lt; 0){
			fprint(2, &#34;ctlproc attached: thread_suspend: %r\n&#34;);
			return -1;
		}

		// Let ptrace tell the process to keep going:
		// then ptrace is out of the way and we&#39;re back in Mach land.
		return ptrace(PT_CONTINUE, id, (caddr_t)1, 0);
	}

	// All the other control messages require a Thread structure.
	if((t = idtotable(id)) == nil){
		werrstr(&#34;no such thread&#34;);
		return -1;
	}

	if(strcmp(msg, &#34;kill&#34;) == 0)
		return ptrace(PT_KILL, t-&gt;pid, 0, 0);

	if(strcmp(msg, &#34;start&#34;) == 0)
		return threadstart(t, 0);

	if(strcmp(msg, &#34;stop&#34;) == 0)
		return threadstop(t);

	if(strcmp(msg, &#34;startstop&#34;) == 0){
		if(threadstart(t, 0) &lt; 0)
			return -1;
		return waitstop(t);
	}

	if(strcmp(msg, &#34;step&#34;) == 0){
		if(threadstart(t, 1) &lt; 0)
			return -1;
		return waitstop(t);
	}

	if(strcmp(msg, &#34;waitstop&#34;) == 0)
		return waitstop(t);

	// sysstop not available on OS X

	werrstr(&#34;unknown control message&#34;);
	return -1;
}

char*
procstatus(int id)
{
	Thread *t;

	if((t = idtotable(id)) == nil)
		return &#34;gone!&#34;;

	if(threadstopped(t))
		return &#34;Stopped&#34;;

	return &#34;Running&#34;;
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

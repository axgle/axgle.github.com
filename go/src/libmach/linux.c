<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/linux.c</title>

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
  <h1 id="generatedHeader">Text file src/libmach/linux.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Plan 9 from User Space src/libmach/Linux.c
// http://code.swtch.com/plan9port/src/tip/src/libmach/Linux.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.
//	Power PC support Copyright © 1995-2004 C H Forsyth (forsyth@terzarima.net).
//	Portions Copyright © 1997-1999 Vita Nuova Limited.
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).
//	Revisions Copyright © 2000-2004 Lucent Technologies Inc. and others.
//	Portions Copyright © 2001-2007 Russ Cox.
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
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

#include &lt;u.h&gt;
#include &lt;sys/syscall.h&gt;	/* for tkill */
#include &lt;unistd.h&gt;
#include &lt;dirent.h&gt;
#include &lt;sys/ptrace.h&gt;
#include &lt;sys/signal.h&gt;
#include &lt;sys/wait.h&gt;
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
#undef waitpid

// The old glibc used with crosstool compilers on thresher
// doesn&#39;t know these numbers, but the Linux kernel
// had them as far back as 2.6.0.
#ifndef WSTOPPED
#define WSTOPPED 2
#define WCONTINUED 8
#define WIFCONTINUED(x) ((x) == 0xffff)
#endif
#ifndef PTRACE_SETOPTIONS
#define PTRACE_SETOPTIONS 0x4200
#define PTRACE_GETEVENTMSG 0x4201
#define PTRACE_O_TRACEFORK 0x2
#define PTRACE_O_TRACEVFORK 0x4
#define PTRACE_O_TRACECLONE 0x8
#define PTRACE_O_TRACEEXEC 0x10
#define PTRACE_O_TRACEVFORKDONE 0x20
#define PTRACE_O_TRACEEXIT 0x40
#define PTRACE_EVENT_FORK 0x1
#define PTRACE_EVENT_VFORK 0x2
#define PTRACE_EVENT_CLONE 0x3
#define PTRACE_EVENT_EXEC 0x4
#define PTRACE_EVENT_VFORK_DONE 0x5
#define PTRACE_EVENT_EXIT 0x6
#endif

typedef struct Ureg64 Ureg64;

static Maprw ptracesegrw;
static Maprw ptraceregrw;

// /usr/include/asm-x86_64/user.h
struct user_regs_struct {
	unsigned long r15,r14,r13,r12,rbp,rbx,r11,r10;
	unsigned long r9,r8,rax,rcx,rdx,rsi,rdi,orig_rax;
	unsigned long rip,cs,eflags;
	unsigned long rsp,ss;
  	unsigned long fs_base, gs_base;
	unsigned long ds,es,fs,gs;
};

// Linux gets very upset if a debugger forgets the reported state
// of a debugged process, so we keep everything we know about
// a debugged process in the LinuxThread structure.
//
// We can poll for state changes by calling waitpid and interpreting
// the integer status code that comes back.  Wait1 does this.
//
// If the process is already running, it is an error to PTRACE_CONT it.
//
// If the process is already stopped, it is an error to stop it again.
//
// If the process is stopped because of a signal, the debugger must
// relay the signal to the PTRACE_CONT call, or else the signal is
// dropped.
//
// If the process exits, the debugger should detach so that the real
// parent can reap the zombie.
//
// On first attach, the debugger should set a handful of flags in order
// to catch future events like fork, clone, exec, etc.

// One for every attached thread.
typedef struct LinuxThread LinuxThread;
struct LinuxThread
{
	int pid;
	int tid;
	int state;
	int signal;
	int child;
	int exitcode;
};

static int trace = 0;

static LinuxThread **thr;
static int nthr;
static int mthr;

static int realpid(int pid);

enum
{
	Unknown,
	Detached,
	Attached,
	AttachStop,
	Stopped,
	Running,
	Forking,
	Vforking,
	VforkDone,
	Cloning,
	Execing,
	Exiting,
	Exited,
	Killed,

	NSTATE,
};

static char* statestr[NSTATE] = {
	&#34;Unknown&#34;,
	&#34;Detached&#34;,
	&#34;Attached&#34;,
	&#34;AttachStop&#34;,
	&#34;Stopped&#34;,
	&#34;Running&#34;,
	&#34;Forking&#34;,
	&#34;Vforking&#34;,
	&#34;VforkDone&#34;,
	&#34;Cloning&#34;,
	&#34;Execing&#34;,
	&#34;Exiting&#34;,
	&#34;Exited&#34;,
	&#34;Killed&#34;
};

static LinuxThread*
attachthread(int pid, int tid, int *new, int newstate)
{
	int i, n, status;
	LinuxThread **p, *t;
	uintptr flags;

	if(new)
		*new = 0;

	for(i=0; i&lt;nthr; i++)
		if((pid == 0 || thr[i]-&gt;pid == pid) &amp;&amp; thr[i]-&gt;tid == tid) {
			t = thr[i];
			goto fixup;
		}

	if(!new)
		return nil;

	if(nthr &gt;= mthr) {
		n = mthr;
		if(n == 0)
			n = 64;
		else
			n *= 2;
		p = realloc(thr, n*sizeof thr[0]);
		if(p == nil)
			return nil;
		thr = p;
		mthr = n;
	}

	t = malloc(sizeof *t);
	if(t == nil)
		return nil;
        memset(t, 0, sizeof *t);

	thr[nthr++] = t;
	t-&gt;pid = pid;
	t-&gt;tid = tid;
	t-&gt;state = newstate;
	if(trace)
		fprint(2, &#34;new thread %d %d\n&#34;, t-&gt;pid, t-&gt;tid);
	if(new)
		*new = 1;

fixup:
	if(t-&gt;state == Detached) {
		if(ptrace(PTRACE_ATTACH, tid, 0, 0) &lt; 0) {
			fprint(2, &#34;ptrace ATTACH %d: %r\n&#34;, tid);
			return nil;
		}
		t-&gt;state = Attached;
	}

	if(t-&gt;state == Attached) {
		// wait for stop, so we can set options
		if(waitpid(tid, &amp;status, __WALL|WUNTRACED|WSTOPPED) &lt; 0)
			return nil;
		if(!WIFSTOPPED(status)) {
			fprint(2, &#34;waitpid %d: status=%#x not stopped\n&#34;, tid);
			return nil;
		}
		t-&gt;state = AttachStop;
	}

	if(t-&gt;state == AttachStop) {
		// set options so we&#39;ll find out about new threads
		flags = PTRACE_O_TRACEFORK |
			PTRACE_O_TRACEVFORK |
			PTRACE_O_TRACECLONE |
			PTRACE_O_TRACEEXEC |
			PTRACE_O_TRACEVFORKDONE |
			PTRACE_O_TRACEEXIT;
		if(ptrace(PTRACE_SETOPTIONS, tid, 0, (void*)flags) &lt; 0)	{
			fprint(2, &#34;ptrace PTRACE_SETOPTIONS %d: %r\n&#34;, tid);
			return nil;
		}
		t-&gt;state = Stopped;
	}

	return t;
}

static LinuxThread*
findthread(int tid)
{
	return attachthread(0, tid, nil, 0);
}

int
procthreadpids(int pid, int *p, int np)
{
	int i, n;
	LinuxThread *t;

	n = 0;
	for(i=0; i&lt;nthr; i++) {
		t = thr[i];
		if(t-&gt;pid == pid) {
			switch(t-&gt;state) {
			case Exited:
			case Detached:
			case Killed:
				break;

			default:
				if(n &lt; np)
					p[n] = t-&gt;tid;
				n++;
				break;
			}
		}
	}
	return n;
}

// Execute a single wait and update the corresponding thread.
static int
wait1(int nohang)
{
	int tid, new, status, event;
	ulong data;
	LinuxThread *t;
	enum
	{
		NormalStop = 0x137f
	};

	if(nohang != 0)
		nohang = WNOHANG;

	tid = waitpid(-1, &amp;status, __WALL|WUNTRACED|WSTOPPED|WCONTINUED|nohang);
	if(tid &lt; 0)
		return -1;
	if(tid == 0)
		return 0;

	if(trace &gt; 0 &amp;&amp; status != NormalStop)
		fprint(2, &#34;TID %d: %#x\n&#34;, tid, status);

	// If we&#39;ve not heard of this tid, something is wrong.
	t = findthread(tid);
	if(t == nil) {
		fprint(2, &#34;ptrace waitpid: unexpected new tid %d status %#x\n&#34;, tid, status);
		return -1;
	}

	if(WIFSTOPPED(status)) {
		t-&gt;state = Stopped;
		t-&gt;signal = WSTOPSIG(status);
		if(trace)
			fprint(2, &#34;tid %d: stopped %#x%s\n&#34;, tid, status,
				status != NormalStop ? &#34; ***&#34; : &#34;&#34;);
		if(t-&gt;signal == SIGTRAP &amp;&amp; (event = status&gt;&gt;16) != 0) {	// ptrace event
			switch(event) {
			case PTRACE_EVENT_FORK:
				t-&gt;state = Forking;
				goto child;

			case PTRACE_EVENT_VFORK:
				t-&gt;state = Vforking;
				goto child;

			case PTRACE_EVENT_CLONE:
				t-&gt;state = Cloning;
				goto child;

			child:
				if(ptrace(PTRACE_GETEVENTMSG, t-&gt;tid, 0, &amp;data) &lt; 0) {
					fprint(2, &#34;ptrace GETEVENTMSG tid %d: %r\n&#34;, tid);
					break;
				}
				t-&gt;child = data;
				attachthread(t-&gt;pid, t-&gt;child, &amp;new, Running);
				if(!new)
					fprint(2, &#34;ptrace child: not new\n&#34;);
				break;

			case PTRACE_EVENT_EXEC:
				t-&gt;state = Execing;
				break;

			case PTRACE_EVENT_VFORK_DONE:
				t-&gt;state = VforkDone;
				break;

			case PTRACE_EVENT_EXIT:
				if(trace)
					fprint(2, &#34;tid %d: exiting %#x\n&#34;, tid, status);
				t-&gt;state = Exiting;
				if(ptrace(PTRACE_GETEVENTMSG, t-&gt;tid, 0, &amp;data) &lt; 0) {
					fprint(2, &#34;ptrace GETEVENTMSG tid %d: %r\n&#34;, tid);
					break;
				}
				t-&gt;exitcode = data;
				break;
			}
		}
	}
	if(WIFCONTINUED(status)) {
		if(trace)
			fprint(2, &#34;tid %d: continued %#x\n&#34;, tid, status);
		t-&gt;state = Running;
	}
	if(WIFEXITED(status)) {
		if(trace)
			fprint(2, &#34;tid %d: exited %#x\n&#34;, tid, status);
		t-&gt;state = Exited;
		t-&gt;exitcode = WEXITSTATUS(status);
		t-&gt;signal = -1;
		ptrace(PTRACE_DETACH, t-&gt;tid, 0, 0);
		if(trace)
			fprint(2, &#34;tid %d: detach exited\n&#34;, tid);
	}
	if(WIFSIGNALED(status)) {
		if(trace)
			fprint(2, &#34;tid %d: signaled %#x\n&#34;, tid, status);
		t-&gt;state = Exited;
		t-&gt;signal = WTERMSIG(status);
		t-&gt;exitcode = -1;
		ptrace(PTRACE_DETACH, t-&gt;tid, 0, 0);
		if(trace)
			fprint(2, &#34;tid %d: detach signaled\n&#34;, tid);
	}
	return 1;
}

static int
waitstop(LinuxThread *t)
{
	while(t-&gt;state == Running)
		if(wait1(0) &lt; 0)
			return -1;
	return 0;
}

// Attach to and stop all threads in process pid.
// Must stop everyone in order to make sure we set
// the &#34;tell me about new threads&#34; option in every
// task.
int
attachallthreads(int pid)
{
	int tid, foundnew, new;
	char buf[100];
	DIR *d;
	struct dirent *de;
	LinuxThread *t;

	if(pid == 0) {
		fprint(2, &#34;attachallthreads(0)\n&#34;);
		return -1;
	}

	pid = realpid(pid);

	snprint(buf, sizeof buf, &#34;/proc/%d/task&#34;, pid);
	if((d = opendir(buf)) == nil) {
		fprint(2, &#34;opendir %s: %r\n&#34;, buf);
		return -1;
	}

	// Loop in case new threads are being created right now.
	// We stop every thread as we find it, so eventually
	// this has to stop (or the system runs out of procs).
	do {
		foundnew = 0;
		while((de = readdir(d)) != nil) {
			tid = atoi(de-&gt;d_name);
			if(tid == 0)
				continue;
			t = attachthread(pid, tid, &amp;new, Detached);
			foundnew |= new;
			if(t)
				waitstop(t);
		}
		rewinddir(d);
	} while(foundnew);
	closedir(d);

	return 0;
}

Map*
attachproc(int pid, Fhdr *fp)
{
	Map *map;

	if(pid == 0) {
		fprint(2, &#34;attachproc(0)\n&#34;);
		return nil;
	}

	if(findthread(pid) == nil &amp;&amp; attachallthreads(pid) &lt; 0)
		return nil;

	map = newmap(0, 4);
	if (!map)
		return 0;
	map-&gt;pid = pid;
	if(mach-&gt;regsize)
		setmap(map, -1, 0, mach-&gt;regsize, 0, &#34;regs&#34;, ptraceregrw);
//	if(mach-&gt;fpregsize)
//		setmap(map, -1, mach-&gt;regsize, mach-&gt;regsize+mach-&gt;fpregsize, 0, &#34;fpregs&#34;, ptraceregrw);
	setmap(map, -1, fp-&gt;txtaddr, fp-&gt;txtaddr+fp-&gt;txtsz, fp-&gt;txtaddr, &#34;*text&#34;, ptracesegrw);
	setmap(map, -1, fp-&gt;dataddr, mach-&gt;utop, fp-&gt;dataddr, &#34;*data&#34;, ptracesegrw);
	return map;
}

void
detachproc(Map *m)
{
	LinuxThread *t;

	t = findthread(m-&gt;pid);
	if(t != nil) {
		ptrace(PTRACE_DETACH, t-&gt;tid, 0, 0);
		t-&gt;state = Detached;
		if(trace)
			fprint(2, &#34;tid %d: detachproc\n&#34;, t-&gt;tid);
		// TODO(rsc): Reclaim thread structs somehow?
	}
	free(m);
}

/* /proc/pid/stat contains
	pid
	command in parens
	0. state
	1. ppid
	2. pgrp
	3. session
	4. tty_nr
	5. tpgid
	6. flags (math=4, traced=10)
	7. minflt
	8. cminflt
	9. majflt
	10. cmajflt
	11. utime
	12. stime
	13. cutime
	14. cstime
	15. priority
	16. nice
	17. 0
	18. itrealvalue
	19. starttime
	20. vsize
	21. rss
	22. rlim
	23. startcode
	24. endcode
	25. startstack
	26. kstkesp
	27. kstkeip
	28. pending signal bitmap
	29. blocked signal bitmap
	30. ignored signal bitmap
	31. caught signal bitmap
	32. wchan
	33. nswap
	34. cnswap
	35. exit_signal
	36. processor
*/

static int
readstat(int pid, char *buf, int nbuf, char **f, int nf)
{
	int fd, n;
	char *p;

	snprint(buf, nbuf, &#34;/proc/%d/stat&#34;, pid);
	if((fd = open(buf, OREAD)) &lt; 0){
		fprint(2, &#34;open %s: %r\n&#34;, buf);
		return -1;
	}
	n = read(fd, buf, nbuf-1);
	close(fd);
	if(n &lt;= 0){
		fprint(2, &#34;read %s: %r\n&#34;, buf);
		return -1;
	}
	buf[n] = 0;

	/* command name is in parens, no parens afterward */
	p = strrchr(buf, &#39;)&#39;);
	if(p == nil || *++p != &#39; &#39;){
		fprint(2, &#34;bad format in /proc/%d/stat\n&#34;, pid);
		return -1;
	}
	++p;

	nf = tokenize(p, f, nf);
	if(0) print(&#34;code 0x%lux-0x%lux stack 0x%lux kstk 0x%lux keip 0x%lux pending 0x%lux\n&#34;,
		strtoul(f[23], 0, 0), strtoul(f[24], 0, 0), strtoul(f[25], 0, 0),
		strtoul(f[26], 0, 0), strtoul(f[27], 0, 0), strtoul(f[28], 0, 0));

	return nf;
}

static char*
readstatus(int pid, char *buf, int nbuf, char *key)
{
	int fd, n;
	char *p;

	snprint(buf, nbuf, &#34;/proc/%d/status&#34;, pid);
	if((fd = open(buf, OREAD)) &lt; 0){
		fprint(2, &#34;open %s: %r\n&#34;, buf);
		return nil;
	}
	n = read(fd, buf, nbuf-1);
	close(fd);
	if(n &lt;= 0){
		fprint(2, &#34;read %s: %r\n&#34;, buf);
		return nil;
	}
	buf[n] = 0;
	p = strstr(buf, key);
	if(p)
		return p+strlen(key);
	return nil;
}

int
procnotes(int pid, char ***pnotes)
{
	char buf[1024], *f[40];
	int i, n, nf;
	char *s, **notes;
	ulong sigs;
	extern char *_p9sigstr(int, char*);

	*pnotes = nil;
	nf = readstat(pid, buf, sizeof buf, f, nelem(f));
	if(nf &lt;= 28)
		return -1;

	sigs = strtoul(f[28], 0, 0) &amp; ~(1&lt;&lt;SIGCONT);
	if(sigs == 0){
		*pnotes = nil;
		return 0;
	}

	notes = malloc(32*sizeof(char*));
	if(notes == nil)
		return -1;
	memset(notes, 0, 32*sizeof(char*));
	n = 0;
	for(i=0; i&lt;32; i++){
		if((sigs&amp;(1&lt;&lt;i)) == 0)
			continue;
		if((s = _p9sigstr(i, nil)) == nil)
			continue;
		notes[n++] = s;
	}
	*pnotes = notes;
	return n;
}

static int
realpid(int pid)
{
	char buf[1024], *p;

	p = readstatus(pid, buf, sizeof buf, &#34;\nTgid:&#34;);
	if(p == nil)
		return pid;
	return atoi(p);
}

int
ctlproc(int pid, char *msg)
{
	int new;
	LinuxThread *t;
	uintptr data;

	while(wait1(1) &gt; 0)
		;

	if(strcmp(msg, &#34;attached&#34;) == 0){
		t = attachthread(pid, pid, &amp;new, Attached);
		if(t == nil)
			return -1;
		return 0;
	}

	if(strcmp(msg, &#34;hang&#34;) == 0){
		if(pid == getpid())
			return ptrace(PTRACE_TRACEME, 0, 0, 0);
		werrstr(&#34;can only hang self&#34;);
		return -1;
	}

	t = findthread(pid);
	if(t == nil) {
		werrstr(&#34;not attached to pid %d&#34;, pid);
		return -1;
	}
	if(t-&gt;state == Exited) {
		werrstr(&#34;pid %d has exited&#34;, pid);
		return -1;
	}
	if(t-&gt;state == Killed) {
		werrstr(&#34;pid %d has been killed&#34;, pid);
		return -1;
	}

	if(strcmp(msg, &#34;kill&#34;) == 0) {
		if(ptrace(PTRACE_KILL, pid, 0, 0) &lt; 0)
			return -1;
		t-&gt;state = Killed;
		return 0;
	}
	if(strcmp(msg, &#34;startstop&#34;) == 0){
		if(ctlproc(pid, &#34;start&#34;) &lt; 0)
			return -1;
		return waitstop(t);
	}
	if(strcmp(msg, &#34;sysstop&#34;) == 0){
		if(ptrace(PTRACE_SYSCALL, pid, 0, 0) &lt; 0)
			return -1;
		t-&gt;state = Running;
		return waitstop(t);
	}
	if(strcmp(msg, &#34;stop&#34;) == 0){
		if(trace &gt; 1)
			fprint(2, &#34;tid %d: tkill stop\n&#34;, pid);
		if(t-&gt;state == Stopped)
			return 0;
		if(syscall(__NR_tkill, pid, SIGSTOP) &lt; 0)
			return -1;
		return waitstop(t);
	}
	if(strcmp(msg, &#34;step&#34;) == 0){
		if(t-&gt;state == Running) {
			werrstr(&#34;cannot single-step unstopped %d&#34;, pid);
			return -1;
		}
		if(ptrace(PTRACE_SINGLESTEP, pid, 0, 0) &lt; 0)
			return -1;
		return waitstop(t);
	}
	if(strcmp(msg, &#34;start&#34;) == 0) {
		if(t-&gt;state == Running)
			return 0;
		data = 0;
		if(t-&gt;state == Stopped &amp;&amp; t-&gt;signal != SIGSTOP &amp;&amp; t-&gt;signal != SIGTRAP)
			data = t-&gt;signal;
		if(trace &amp;&amp; data)
			fprint(2, &#34;tid %d: continue %lud\n&#34;, pid, (ulong)data);
		if(ptrace(PTRACE_CONT, pid, 0, (void*)data) &lt; 0)
			return -1;
		t-&gt;state = Running;
		return 0;
	}
	if(strcmp(msg, &#34;waitstop&#34;) == 0) {
		return waitstop(t);
	}
	werrstr(&#34;unknown control message &#39;%s&#39;&#34;, msg);
	return -1;
}

char*
proctextfile(int pid)
{
	static char buf[1024], pbuf[128];

	snprint(pbuf, sizeof pbuf, &#34;/proc/%d/exe&#34;, pid);
	if(readlink(pbuf, buf, sizeof buf) &gt;= 0)
		return strdup(buf);
	if(access(pbuf, AEXIST) &gt;= 0)
		return strdup(pbuf);
	return nil;
}


static int
ptracerw(int type, int xtype, int isr, int pid, uvlong addr, void *v, uint n)
{
	int i;
	uintptr u;
	uchar buf[sizeof(uintptr)];

	for(i=0; i&lt;n; i+=sizeof(uintptr)){
		if(isr){
			errno = 0;
			u = ptrace(type, pid, addr+i, 0);
			if(errno)
				goto ptraceerr;
			if(n-i &gt;= sizeof(uintptr))
				*(uintptr*)((char*)v+i) = u;
			else{
				*(uintptr*)buf = u;
				memmove((char*)v+i, buf, n-i);
			}
		}else{
			if(n-i &gt;= sizeof(uintptr))
				u = *(uintptr*)((char*)v+i);
			else{
				errno = 0;
				u = ptrace(xtype, pid, addr+i, 0);
				if(errno)
					return -1;
				*(uintptr*)buf = u;
				memmove(buf, (char*)v+i, n-i);
				u = *(uintptr*)buf;
			}
			if(ptrace(type, pid, addr+i, u) &lt; 0)
				goto ptraceerr;
		}
	}
	return 0;

ptraceerr:
	werrstr(&#34;ptrace %s addr=%#llux pid=%d: %r&#34;, isr ? &#34;read&#34; : &#34;write&#34;, addr, pid);
	return -1;
}

static int
ptracesegrw(Map *map, Seg *seg, uvlong addr, void *v, uint n, int isr)
{
	return ptracerw(isr ? PTRACE_PEEKDATA : PTRACE_POKEDATA, PTRACE_PEEKDATA,
		isr, map-&gt;pid, addr, v, n);
}

// If the debugger is compiled as an x86-64 program,
// then all the ptrace register read/writes are done on
// a 64-bit register set.  If the target program
// is a 32-bit program, the debugger is expected to
// read the bottom half of the relevant registers
// out of the 64-bit set.

// Linux 32-bit is
//	BX CX DX SI DI BP AX DS ES FS GS OrigAX IP CS EFLAGS SP SS

// Linux 64-bit is
//	R15 R14 R13 R12 BP BX R11 R10 R9 R8 AX CX DX SI DI OrigAX IP CS EFLAGS SP SS FSBase GSBase DS ES FS GS

// Go 32-bit is
//	DI SI BP NSP BX DX CX AX GS FS ES DS TRAP ECODE PC CS EFLAGS SP SS

// uint go32tolinux32tab[] = {
//	4, 3, 5, 15, 0, 2, 1, 6, 10, 9, 8, 7, -1, -1, 12, 13, 14, 15, 16
// };
uint go32tolinux64tab[] = {
	14, 13, 4, 19, 5, 12, 11, 10, 26, 25, 24, 23, -1, -1, 16, 17, 18, 19, 20
};
static int
go32tolinux64(uvlong addr)
{
	int r;

	if(addr%4 || addr/4 &gt;= nelem(go32tolinux64tab))
		return -1;
	r = go32tolinux64tab[addr/4];
	if(r &lt; 0)
		return -1;
	return r*8;
}

extern Mach mi386;

static int
go2linux(uvlong addr)
{
	// TODO(rsc): If this file is being compiled in 32-bit mode,
	// need to use the go32tolinux32 table instead.

	if(mach == &amp;mi386)
		return go32tolinux64(addr);

	switch(addr){
	case offsetof(Ureg64, ax):
		return offsetof(struct user_regs_struct, rax);
	case offsetof(Ureg64, bx):
		return offsetof(struct user_regs_struct, rbx);
	case offsetof(Ureg64, cx):
		return offsetof(struct user_regs_struct, rcx);
	case offsetof(Ureg64, dx):
		return offsetof(struct user_regs_struct, rdx);
	case offsetof(Ureg64, si):
		return offsetof(struct user_regs_struct, rsi);
	case offsetof(Ureg64, di):
		return offsetof(struct user_regs_struct, rdi);
	case offsetof(Ureg64, bp):
		return offsetof(struct user_regs_struct, rbp);
	case offsetof(Ureg64, r8):
		return offsetof(struct user_regs_struct, r8);
	case offsetof(Ureg64, r9):
		return offsetof(struct user_regs_struct, r9);
	case offsetof(Ureg64, r10):
		return offsetof(struct user_regs_struct, r10);
	case offsetof(Ureg64, r11):
		return offsetof(struct user_regs_struct, r11);
	case offsetof(Ureg64, r12):
		return offsetof(struct user_regs_struct, r12);
	case offsetof(Ureg64, r13):
		return offsetof(struct user_regs_struct, r13);
	case offsetof(Ureg64, r14):
		return offsetof(struct user_regs_struct, r14);
	case offsetof(Ureg64, r15):
		return offsetof(struct user_regs_struct, r15);
	case offsetof(Ureg64, ds):
		return offsetof(struct user_regs_struct, ds);
	case offsetof(Ureg64, es):
		return offsetof(struct user_regs_struct, es);
	case offsetof(Ureg64, fs):
		return offsetof(struct user_regs_struct, fs);
	case offsetof(Ureg64, gs):
		return offsetof(struct user_regs_struct, gs);
	case offsetof(Ureg64, ip):
		return offsetof(struct user_regs_struct, rip);
	case offsetof(Ureg64, cs):
		return offsetof(struct user_regs_struct, cs);
	case offsetof(Ureg64, flags):
		return offsetof(struct user_regs_struct, eflags);
	case offsetof(Ureg64, sp):
		return offsetof(struct user_regs_struct, rsp);
	case offsetof(Ureg64, ss):
		return offsetof(struct user_regs_struct, ss);
	}
	return -1;
}

static int
ptraceregrw(Map *map, Seg *seg, uvlong addr, void *v, uint n, int isr)
{
	int laddr;
	uvlong u;

	if((laddr = go2linux(addr)) &lt; 0){
		if(isr){
			memset(v, 0, n);
			return 0;
		}
		werrstr(&#34;register %llud not available&#34;, addr);
		return -1;
	}

	if(isr){
		errno = 0;
		u = ptrace(PTRACE_PEEKUSER, map-&gt;pid, laddr, 0);
		if(errno)
			goto ptraceerr;
		switch(n){
		case 1:
			*(uint8*)v = u;
			break;
		case 2:
			*(uint16*)v = u;
			break;
		case 4:
			*(uint32*)v = u;
			break;
		case 8:
			*(uint64*)v = u;
			break;
		default:
			werrstr(&#34;bad register size&#34;);
			return -1;
		}
	}else{
		switch(n){
		case 1:
			u = *(uint8*)v;
			break;
		case 2:
			u = *(uint16*)v;
			break;
		case 4:
			u = *(uint32*)v;
			break;
		case 8:
			u = *(uint64*)v;
			break;
		default:
			werrstr(&#34;bad register size&#34;);
			return -1;
		}
		if(ptrace(PTRACE_POKEUSER, map-&gt;pid, laddr, (void*)(uintptr)u) &lt; 0)
			goto ptraceerr;
	}
	return 0;

ptraceerr:
	werrstr(&#34;ptrace %s register laddr=%d pid=%d n=%d: %r&#34;, isr ? &#34;read&#34; : &#34;write&#34;, laddr, map-&gt;pid, n);
	return -1;
}

char*
procstatus(int pid)
{
	LinuxThread *t;

	t = findthread(pid);
	if(t == nil)
		return &#34;???&#34;;

	return statestr[t-&gt;state];
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/prof/main.c</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

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
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Text file src/cmd/prof/main.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &lt;u.h&gt;
#include &lt;time.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;ctype.h&gt;

#include &lt;ureg_amd64.h&gt;
#include &lt;mach.h&gt;

char* file = &#34;6.out&#34;;
static Fhdr fhdr;
int have_syms;
int fd;
Map	*symmap;
struct Ureg ureg;
int total_sec = 0;
int delta_msec = 100;
int nsample;
int nsamplethread;

// output formats
int functions;	// print functions
int histograms;	// print histograms
int linenums;	// print file and line numbers rather than function names
int registers;	// print registers
int stacks;		// print stack traces

int pid;		// main process pid

int nthread;	// number of threads
int thread[32];	// thread pids
Map *map[32];	// thread maps

void
Usage(void)
{
	fprint(2, &#34;Usage: prof -p pid [-t total_secs] [-d delta_msec] [6.out args ...]\n&#34;);
	fprint(2, &#34;\tformats (default -h):\n&#34;);
	fprint(2, &#34;\t\t-h: histograms\n&#34;);
	fprint(2, &#34;\t\t-f: dynamic functions\n&#34;);
	fprint(2, &#34;\t\t-l: dynamic file and line numbers\n&#34;);
	fprint(2, &#34;\t\t-r: dynamic registers\n&#34;);
	fprint(2, &#34;\t\t-s: dynamic function stack traces\n&#34;);
	fprint(2, &#34;\t\t-hs: include stack info in histograms\n&#34;);
	exit(2);
}

typedef struct PC PC;
struct PC {
	uvlong pc;
	uvlong callerpc;
	unsigned int count;
	PC* next;
};

enum {
	Ncounters = 256
};

PC *counters[Ncounters];

void
regprint(void)
{
	fprint(2, &#34;ax\t0x%llux\n&#34;, ureg.ax);
	fprint(2, &#34;bx\t0x%llux\n&#34;, ureg.bx);
	fprint(2, &#34;cx\t0x%llux\n&#34;, ureg.cx);
	fprint(2, &#34;dx\t0x%llux\n&#34;, ureg.dx);
	fprint(2, &#34;si\t0x%llux\n&#34;, ureg.si);
	fprint(2, &#34;di\t0x%llux\n&#34;, ureg.di);
	fprint(2, &#34;bp\t0x%llux\n&#34;, ureg.bp);
	fprint(2, &#34;r8\t0x%llux\n&#34;, ureg.r8);
	fprint(2, &#34;r9\t0x%llux\n&#34;, ureg.r9);
	fprint(2, &#34;r10\t0x%llux\n&#34;, ureg.r10);
	fprint(2, &#34;r11\t0x%llux\n&#34;, ureg.r11);
	fprint(2, &#34;r12\t0x%llux\n&#34;, ureg.r12);
	fprint(2, &#34;r13\t0x%llux\n&#34;, ureg.r13);
	fprint(2, &#34;r14\t0x%llux\n&#34;, ureg.r14);
	fprint(2, &#34;r15\t0x%llux\n&#34;, ureg.r15);
	fprint(2, &#34;ds\t0x%llux\n&#34;, ureg.ds);
	fprint(2, &#34;es\t0x%llux\n&#34;, ureg.es);
	fprint(2, &#34;fs\t0x%llux\n&#34;, ureg.fs);
	fprint(2, &#34;gs\t0x%llux\n&#34;, ureg.gs);
	fprint(2, &#34;type\t0x%llux\n&#34;, ureg.type);
	fprint(2, &#34;error\t0x%llux\n&#34;, ureg.error);
	fprint(2, &#34;pc\t0x%llux\n&#34;, ureg.ip);
	fprint(2, &#34;cs\t0x%llux\n&#34;, ureg.cs);
	fprint(2, &#34;flags\t0x%llux\n&#34;, ureg.flags);
	fprint(2, &#34;sp\t0x%llux\n&#34;, ureg.sp);
	fprint(2, &#34;ss\t0x%llux\n&#34;, ureg.ss);
}

int
getthreads(void)
{
	int i, j, curn, found;
	Map *curmap[nelem(map)];
	int curthread[nelem(map)];
	static int complained = 0;

	curn = procthreadpids(pid, curthread, nelem(curthread));
	if(curn &lt;= 0)
		return curn;

	if(curn &gt; nelem(map)) {
		if(complained == 0) {
			fprint(2, &#34;prof: too many threads; limiting to %d\n&#34;, nthread, nelem(map));
			complained = 1;
		}
		curn = nelem(map);
	}
	if(curn == nthread &amp;&amp; memcmp(thread, curthread, curn*sizeof(*thread)) == 0)
		return curn;	// no changes

	// Number of threads has changed (might be the init case).
	// A bit expensive but rare enough not to bother being clever.
	for(i = 0; i &lt; curn; i++) {
		found = 0;
		for(j = 0; j &lt; nthread; j++) {
			if(curthread[i] == thread[j]) {
				found = 1;
				curmap[i] = map[j];
				map[j] = nil;
				break;
			}
		}
		if(found)
			continue;

		// map new thread
		curmap[i] = attachproc(curthread[i], &amp;fhdr);
		if(curmap[i] == nil) {
			fprint(2, &#34;prof: can&#39;t attach to %d: %r\n&#34;, curthread[i]);
			return -1;
		}
	}

	for(j = 0; j &lt; nthread; j++)
		if(map[j] != nil)
			detachproc(map[j]);

	nthread = curn;
	memmove(thread, curthread, nthread*sizeof thread[0]);
	memmove(map, curmap, sizeof map);
	return nthread;
}

int
sample(Map *map)
{
	int i;
	static int n;

	n++;
	if(registers) {
		for(i = 0; i &lt; sizeof ureg; i+=8) {
			if(get8(map, (uvlong)i, &amp;((uvlong*)&amp;ureg)[i/8]) &lt; 0)
				goto bad;
		}
	} else {
		// we need only two registers
		if(get8(map, offsetof(struct Ureg, ip), (uvlong*)&amp;ureg.ip) &lt; 0)
			goto bad;
		if(get8(map, offsetof(struct Ureg, sp), (uvlong*)&amp;ureg.sp) &lt; 0)
			goto bad;
	}
	return 1;
bad:
	if(n == 1)
		fprint(2, &#34;prof: can&#39;t read registers: %r\n&#34;);
	return 0;
}

void
addtohistogram(uvlong pc, uvlong callerpc, uvlong sp)
{
	int h;
	PC *x;

	h = (pc + callerpc*101) % Ncounters;
	for(x = counters[h]; x != NULL; x = x-&gt;next) {
		if(x-&gt;pc == pc &amp;&amp; x-&gt;callerpc == callerpc) {
			x-&gt;count++;
			return;
		}
	}
	x = malloc(sizeof(PC));
	x-&gt;pc = pc;
	x-&gt;callerpc = callerpc;
	x-&gt;count = 1;
	x-&gt;next = counters[h];
	counters[h] = x;
}

uvlong nextpc;

void
xptrace(Map *map, uvlong pc, uvlong sp, Symbol *sym)
{
	char buf[1024];
	if(sym == nil){
		fprint(2, &#34;syms\n&#34;);
		return;
	}
	if(histograms)
		addtohistogram(nextpc, pc, sp);
	if(!histograms || stacks &gt; 1) {
		if(nextpc == 0)
			nextpc = sym-&gt;value;
		fprint(2, &#34;%s(&#34;, sym-&gt;name);
		fprint(2, &#34;)&#34;);
		if(nextpc != sym-&gt;value)
			fprint(2, &#34;+%#llux &#34;, nextpc - sym-&gt;value);
		if(have_syms &amp;&amp; linenums &amp;&amp; fileline(buf, sizeof buf, pc)) {
			fprint(2, &#34; %s&#34;, buf);
		}
		fprint(2, &#34;\n&#34;);
	}
	nextpc = pc;
}

void
stacktracepcsp(Map *map, uvlong pc, uvlong sp)
{
	nextpc = pc;
	if(machdata-&gt;ctrace==nil)
		fprint(2, &#34;no machdata-&gt;ctrace\n&#34;);
	else if(machdata-&gt;ctrace(map, pc, sp, 0, xptrace) &lt;= 0)
		fprint(2, &#34;no stack frame: pc=%#p sp=%#p\n&#34;, pc, sp);
	else {
		addtohistogram(nextpc, 0, sp);
		if(!histograms || stacks &gt; 1)
			fprint(2, &#34;\n&#34;);
	}
}

void
printpc(Map *map, uvlong pc, uvlong sp)
{
	char buf[1024];
	if(registers)
		regprint();
	if(have_syms &gt; 0 &amp;&amp; linenums &amp;&amp;  fileline(buf, sizeof buf, pc))
		fprint(2, &#34;%s\n&#34;, buf);
	if(have_syms &gt; 0 &amp;&amp; functions) {
		symoff(buf, sizeof(buf), pc, CANY);
		fprint(2, &#34;%s\n&#34;, buf);
	}
	if(stacks){
		stacktracepcsp(map, pc, sp);
	}
	else if(histograms){
		addtohistogram(pc, 0, sp);
	}
}

void
samples(void)
{
	int i, pid, msec;
	struct timespec req;

	req.tv_sec = delta_msec/1000;
	req.tv_nsec = 1000000*(delta_msec % 1000);
	for(msec = 0; total_sec &lt;= 0 || msec &lt; 1000*total_sec; msec += delta_msec) {
		nsample++;
		nsamplethread += nthread;
		for(i = 0; i &lt; nthread; i++) {
			pid = thread[i];
			if(ctlproc(pid, &#34;stop&#34;) &lt; 0)
				return;
			if(!sample(map[i])) {
				ctlproc(pid, &#34;start&#34;);
				return;
			}
			printpc(map[i], ureg.ip, ureg.sp);
			ctlproc(pid, &#34;start&#34;);
		}
		nanosleep(&amp;req, NULL);
		getthreads();
		if(nthread == 0)
			break;
	}
}

typedef struct Func Func;
struct Func
{
	Func *next;
	Symbol s;
	uint onstack;
	uint leaf;
};

Func *func[257];
int nfunc;

Func*
findfunc(uvlong pc)
{
	Func *f;
	uint h;
	Symbol s;

	if(pc == 0)
		return nil;

	if(!findsym(pc, CTEXT, &amp;s))
		return nil;

	h = s.value % nelem(func);
	for(f = func[h]; f != NULL; f = f-&gt;next)
		if(f-&gt;s.value == s.value)
			return f;

	f = malloc(sizeof *f);
	memset(f, 0, sizeof *f);
	f-&gt;s = s;
	f-&gt;next = func[h];
	func[h] = f;
	nfunc++;
	return f;
}

int
compareleaf(const void *va, const void *vb)
{
	Func *a, *b;

	a = *(Func**)va;
	b = *(Func**)vb;
	if(a-&gt;leaf != b-&gt;leaf)
		return b-&gt;leaf - a-&gt;leaf;
	if(a-&gt;onstack != b-&gt;onstack)
		return b-&gt;onstack - a-&gt;onstack;
	return strcmp(a-&gt;s.name, b-&gt;s.name);
}

void
dumphistogram()
{
	int i, h, n;
	PC *x;
	Func *f, **ff;

	if(!histograms)
		return;

	// assign counts to functions.
	for(h = 0; h &lt; Ncounters; h++) {
		for(x = counters[h]; x != NULL; x = x-&gt;next) {
			f = findfunc(x-&gt;pc);
			if(f) {
				f-&gt;onstack += x-&gt;count;
				f-&gt;leaf += x-&gt;count;
			}
			f = findfunc(x-&gt;callerpc);
			if(f)
				f-&gt;leaf -= x-&gt;count;
		}
	}

	// build array
	ff = malloc(nfunc*sizeof ff[0]);
	n = 0;
	for(h = 0; h &lt; nelem(func); h++)
		for(f = func[h]; f != NULL; f = f-&gt;next)
			ff[n++] = f;

	// sort by leaf counts
	qsort(ff, nfunc, sizeof ff[0], compareleaf);

	// print.
	fprint(2, &#34;%d samples (avg %.1g threads)\n&#34;, nsample, (double)nsamplethread/nsample);
	for(i = 0; i &lt; nfunc; i++) {
		f = ff[i];
		fprint(2, &#34;%6.2f%%\t&#34;, 100.0*(double)f-&gt;leaf/nsample);
		if(stacks)
			fprint(2, &#34;%6.2f%%\t&#34;, 100.0*(double)f-&gt;onstack/nsample);
		fprint(2, &#34;%s\n&#34;, f-&gt;s.name);
	}
}

int
startprocess(char **argv)
{
	int pid;

	if((pid = fork()) == 0) {
		pid = getpid();
		if(ctlproc(pid, &#34;hang&#34;) &lt; 0){
			fprint(2, &#34;prof: child process could not hang\n&#34;);
			exits(0);
		}
		execv(argv[0], argv);
		fprint(2, &#34;prof: could not exec %s: %r\n&#34;, argv[0]);
		exits(0);
	}

	if(pid == -1) {
		fprint(2, &#34;prof: could not fork\n&#34;);
		exit(1);
	}
	if(ctlproc(pid, &#34;attached&#34;) &lt; 0 || ctlproc(pid, &#34;waitstop&#34;) &lt; 0) {
		fprint(2, &#34;prof: could not attach to child process: %r\n&#34;);
		exit(1);
	}
	return pid;
}

void
detach(void)
{
	int i;

	for(i = 0; i &lt; nthread; i++)
		detachproc(map[i]);
}

int
main(int argc, char *argv[])
{
	int i;

	ARGBEGIN{
	case &#39;d&#39;:
		delta_msec = atoi(EARGF(Usage()));
		break;
	case &#39;t&#39;:
		total_sec = atoi(EARGF(Usage()));
		break;
	case &#39;p&#39;:
		pid = atoi(EARGF(Usage()));
		break;
	case &#39;f&#39;:
		functions = 1;
		break;
	case &#39;h&#39;:
		histograms = 1;
		break;
	case &#39;l&#39;:
		linenums = 1;
		break;
	case &#39;r&#39;:
		registers = 1;
		break;
	case &#39;s&#39;:
		stacks++;
		break;
	}ARGEND
	if(pid &lt;= 0 &amp;&amp; argc == 0)
		Usage();
	if(functions+linenums+registers+stacks == 0)
		histograms = 1;
	if(!machbyname(&#34;amd64&#34;)) {
		fprint(2, &#34;prof: no amd64 support\n&#34;, pid);
		exit(1);
	}
	if(argc &gt; 0)
		file = argv[0];
	else if(pid)
		file = proctextfile(pid);
	fd = open(file, 0);
	if(fd &lt; 0) {
		fprint(2, &#34;prof: can&#39;t open %s: %r\n&#34;, file);
		exit(1);
	}
	if(crackhdr(fd, &amp;fhdr)) {
		have_syms = syminit(fd, &amp;fhdr);
		if(!have_syms) {
			fprint(2, &#34;prof: no symbols for %s: %r\n&#34;, file);
		}
	} else {
		fprint(2, &#34;prof: crack header for %s: %r\n&#34;, file);
		exit(1);
	}
	if(pid &lt;= 0)
		pid = startprocess(argv);
	attachproc(pid, &amp;fhdr);	// initializes thread list
	if(getthreads() &lt;= 0) {
		detach();
		fprint(2, &#34;prof: can&#39;t find threads for pid %d\n&#34;, pid);
		exit(1);
	}
	for(i = 0; i &lt; nthread; i++)
		ctlproc(thread[i], &#34;start&#34;);
	samples();
	detach();
	dumphistogram();
	exit(0);
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

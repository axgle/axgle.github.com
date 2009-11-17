<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cov/main.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/cov/main.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * code coverage
 */

#include &lt;u.h&gt;
#include &lt;time.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;ctype.h&gt;
#include &#34;tree.h&#34;

#include &lt;ureg_amd64.h&gt;
#include &lt;mach.h&gt;
typedef struct Ureg Ureg;

void
usage(void)
{
	fprint(2, &#34;usage: cov [-lsv] [-g substring] [-m minlines] [6.out args...]\n&#34;);
	fprint(2, &#34;-g specifies pattern of interesting functions or files\n&#34;);
	exits(&#34;usage&#34;);
}

typedef struct Range Range;
struct Range
{
	uvlong pc;
	uvlong epc;
};

int chatty;
int fd;
int longnames;
int pid;
int doshowsrc;
Map *mem;
Map *text;
Fhdr fhdr;
char *substring;
char cwd[1000];
int ncwd;
int minlines = -1000;

Tree breakpoints;	// code ranges not run

/*
 * comparison for Range structures
 * they are &#34;equal&#34; if they overlap, so
 * that a search for [pc, pc+1) finds the
 * Range containing pc.
 */
int
rangecmp(void *va, void *vb)
{
	Range *a = va, *b = vb;
	if(a-&gt;epc &lt;= b-&gt;pc)
		return 1;
	if(b-&gt;epc &lt;= a-&gt;pc)
		return -1;
	return 0;
}

/*
 * remember that we ran the section of code [pc, epc).
 */
void
ran(uvlong pc, uvlong epc)
{
	Range key;
	Range *r;
	uvlong oldepc;

	if(chatty)
		print(&#34;run %#llux-%#llux\n&#34;, pc, epc);

	key.pc = pc;
	key.epc = pc+1;
	r = treeget(&amp;breakpoints, &amp;key);
	if(r == nil)
		sysfatal(&#34;unchecked breakpoint at %#lux+%d&#34;, pc, (int)(epc-pc));

	// Might be that the tail of the sequence
	// was run already, so r-&gt;epc is before the end.
	// Adjust len.
	if(epc &gt; r-&gt;epc)
		epc = r-&gt;epc;

	if(r-&gt;pc == pc) {
		r-&gt;pc = epc;
	} else {
		// Chop r to before pc;
		// add new entry for after if needed.
		// Changing r-&gt;epc does not affect r&#39;s position in the tree.
		oldepc = r-&gt;epc;
		r-&gt;epc = pc;
		if(epc &lt; oldepc) {
			Range *n;
			n = malloc(sizeof *n);
			n-&gt;pc = epc;
			n-&gt;epc = oldepc;
			treeput(&amp;breakpoints, n, n);
		}
	}
}

void
showsrc(char *file, int line1, int line2)
{
	Biobuf *b;
	char *p;
	int n, stop;

	if((b = Bopen(file, OREAD)) == nil) {
		print(&#34;\topen %s: %r\n&#34;, file);
		return;
	}

	for(n=1; n&lt;line1 &amp;&amp; (p = Brdstr(b, &#39;\n&#39;, 1)) != nil; n++)
		free(p);

	// print up to five lines (this one and 4 more).
	// if there are more than five lines, print 4 and &#34;...&#34;
	stop = n+4;
	if(stop &gt; line2)
		stop = line2;
	if(stop &lt; line2)
		stop--;
	for(; n&lt;=stop &amp;&amp; (p = Brdstr(b, &#39;\n&#39;, 1)) != nil; n++) {
		print(&#34;  %d %s\n&#34;, n, p);
		free(p);
	}
	if(n &lt; line2)
		print(&#34;  ...\n&#34;);
	Bterm(b);
}

/*
 * if s is in the current directory or below,
 * return the relative path.
 */
char*
shortname(char *s)
{
	if(!longnames &amp;&amp; strlen(s) &gt; ncwd &amp;&amp; memcmp(s, cwd, ncwd) == 0 &amp;&amp; s[ncwd] == &#39;/&#39;)
		return s+ncwd+1;
	return s;
}

/*
 * we&#39;ve decided that [pc, epc) did not run.
 * do something about it.
 */
void
missing(uvlong pc, uvlong epc)
{
	char file[1000];
	int line1, line2;
	char buf[100];
	Symbol s;
	char *p;
	uvlong uv;

	if(!findsym(pc, CTEXT, &amp;s) || !fileline(file, sizeof file, pc)) {
	notfound:
		print(&#34;%#llux-%#llux\n&#34;, pc, epc);
		return;
	}
	p = strrchr(file, &#39;:&#39;);
	*p++ = 0;
	line1 = atoi(p);
	for(uv=pc; uv&lt;epc; ) {
		if(!fileline(file, sizeof file, epc-2))
			goto notfound;
		uv += machdata-&gt;instsize(text, uv);
	}
	p = strrchr(file, &#39;:&#39;);
	*p++ = 0;
	line2 = atoi(p);

	if(line2+1-line2 &lt; minlines)
		return;

	if(pc == s.value) {
		// never entered function
		print(&#34;%s:%d %s never called (%#llux-%#llux)\n&#34;, shortname(file), line1, s.name, pc, epc);
		return;
	}
	if(pc &lt;= s.value+13) {
		// probably stub for stack growth.
		// check whether last instruction is call to morestack.
		// the -5 below is the length of
		//	CALL sys.morestack.
		buf[0] = 0;
		machdata-&gt;das(text, epc-5, 0, buf, sizeof buf);
		if(strstr(buf, &#34;morestack&#34;))
			return;
	}

	if(epc - pc == 5) {
		// check for CALL sys.throwindex
		buf[0] = 0;
		machdata-&gt;das(text, pc, 0, buf, sizeof buf);
		if(strstr(buf, &#34;throwindex&#34;))
			return;
	}

	if(epc - pc == 2 || epc -pc == 3) {
		// check for XORL inside shift.
		// (on x86 have to implement large left or unsigned right shift with explicit zeroing).
		//	f+90 0x00002c9f	CMPL	CX,$20
		//	f+93 0x00002ca2	JCS	f+97(SB)
		//	f+95 0x00002ca4	XORL	AX,AX &lt;&lt;&lt;
		//	f+97 0x00002ca6	SHLL	CL,AX
		//	f+99 0x00002ca8	MOVL	$1,CX
		//
		//	f+c8 0x00002cd7	CMPL	CX,$40
		//	f+cb 0x00002cda	JCS	f+d0(SB)
		//	f+cd 0x00002cdc	XORQ	AX,AX &lt;&lt;&lt;
		//	f+d0 0x00002cdf	SHLQ	CL,AX
		//	f+d3 0x00002ce2	MOVQ	$1,CX
		buf[0] = 0;
		machdata-&gt;das(text, pc, 0, buf, sizeof buf);
		if(strncmp(buf, &#34;XOR&#34;, 3) == 0) {
			machdata-&gt;das(text, epc, 0, buf, sizeof buf);
			if(strncmp(buf, &#34;SHL&#34;, 3) == 0 || strncmp(buf, &#34;SHR&#34;, 3) == 0)
				return;
		}
	}

	if(epc - pc == 3) {
		// check for SAR inside shift.
		// (on x86 have to implement large signed right shift as &gt;&gt;31).
		//	f+36 0x00016216	CMPL	CX,$20
		//	f+39 0x00016219	JCS	f+3e(SB)
		//	f+3b 0x0001621b	SARL	$1f,AX &lt;&lt;&lt;
		//	f+3e 0x0001621e	SARL	CL,AX
		//	f+40 0x00016220	XORL	CX,CX
		//	f+42 0x00016222	CMPL	CX,AX
		buf[0] = 0;
		machdata-&gt;das(text, pc, 0, buf, sizeof buf);
		if(strncmp(buf, &#34;SAR&#34;, 3) == 0) {
			machdata-&gt;das(text, epc, 0, buf, sizeof buf);
			if(strncmp(buf, &#34;SAR&#34;, 3) == 0)
				return;
		}
	}

	// show first instruction to make clear where we were.
	machdata-&gt;das(text, pc, 0, buf, sizeof buf);

	if(line1 != line2)
		print(&#34;%s:%d,%d %#llux-%#llux %s\n&#34;,
			shortname(file), line1, line2, pc, epc, buf);
	else
		print(&#34;%s:%d %#llux-%#llux %s\n&#34;,
			shortname(file), line1, pc, epc, buf);
	if(doshowsrc)
		showsrc(file, line1, line2);
}

/*
 * walk the tree, calling missing for each non-empty
 * section of missing code.
 */
void
walktree(TreeNode *t)
{
	Range *n;

	if(t == nil)
		return;
	walktree(t-&gt;left);
	n = t-&gt;key;
	if(n-&gt;pc &lt; n-&gt;epc)
		missing(n-&gt;pc, n-&gt;epc);
	walktree(t-&gt;right);
}

/*
 * set a breakpoint all over [pc, epc)
 * and remember that we did.
 */
void
breakpoint(uvlong pc, uvlong epc)
{
	Range *r;

	r = malloc(sizeof *r);
	r-&gt;pc = pc;
	r-&gt;epc = epc;
	treeput(&amp;breakpoints, r, r);

	for(; pc &lt; epc; pc+=machdata-&gt;bpsize)
		put1(mem, pc, machdata-&gt;bpinst, machdata-&gt;bpsize);
}

/*
 * install breakpoints over all text symbols
 * that match the pattern.
 */
void
cover(void)
{
	Symbol s;
	char *lastfn;
	uvlong lastpc;
	int i;
	char buf[200];

	lastfn = nil;
	lastpc = 0;
	for(i=0; textsym(&amp;s, i); i++) {
		switch(s.type) {
		case &#39;T&#39;:
		case &#39;t&#39;:
			if(lastpc != 0) {
				breakpoint(lastpc, s.value);
				lastpc = 0;
			}
			// Ignore second entry for a given name;
			// that&#39;s the debugging blob.
			if(lastfn &amp;&amp; strcmp(s.name, lastfn) == 0)
				break;
			lastfn = s.name;
			buf[0] = 0;
			fileline(buf, sizeof buf, s.value);
			if(substring == nil || strstr(buf, substring) || strstr(s.name, substring))
				lastpc = s.value;
		}
	}
}

uvlong
rgetzero(Map *map, char *reg)
{
	return 0;
}

/*
 * remove the breakpoints at pc and successive instructions,
 * up to and including the first jump or other control flow transfer.
 */
void
uncover(uvlong pc)
{
	uchar buf[1000];
	int n, n1, n2;
	uvlong foll[2];

	// Double-check that we stopped at a breakpoint.
	if(get1(mem, pc, buf, machdata-&gt;bpsize) &lt; 0)
		sysfatal(&#34;read mem inst at %#llux: %r&#34;, pc);
	if(memcmp(buf, machdata-&gt;bpinst, machdata-&gt;bpsize) != 0)
		sysfatal(&#34;stopped at %#llux; not at breakpoint %d&#34;, pc, machdata-&gt;bpsize);

	// Figure out how many bytes of straight-line code
	// there are in the text starting at pc.
	n = 0;
	while(n &lt; sizeof buf) {
		n1 = machdata-&gt;instsize(text, pc+n);
		if(n+n1 &gt; sizeof buf)
			break;
		n2 = machdata-&gt;foll(text, pc+n, rgetzero, foll);
		n += n1;
		if(n2 != 1 || foll[0] != pc+n)
			break;
	}

	// Record that this section of code ran.
	ran(pc, pc+n);

	// Put original instructions back.
	if(get1(text, pc, buf, n) &lt; 0)
		sysfatal(&#34;get1: %r&#34;);
	if(put1(mem, pc, buf, n) &lt; 0)
		sysfatal(&#34;put1: %r&#34;);
}

int
startprocess(char **argv)
{
	int pid;

	if((pid = fork()) &lt; 0)
		sysfatal(&#34;fork: %r&#34;);
	if(pid == 0) {
		pid = getpid();
		if(ctlproc(pid, &#34;hang&#34;) &lt; 0)
			sysfatal(&#34;ctlproc hang: %r&#34;);
		execv(argv[0], argv);
		sysfatal(&#34;exec %s: %r&#34;, argv[0]);
	}
	if(ctlproc(pid, &#34;attached&#34;) &lt; 0 || ctlproc(pid, &#34;waitstop&#34;) &lt; 0)
		sysfatal(&#34;attach %d %s: %r&#34;, pid, argv[0]);
	return pid;
}

int
go(void)
{
	uvlong pc;
	char buf[100];
	int n;

	for(n = 0;; n++) {
		ctlproc(pid, &#34;startstop&#34;);
		if(get8(mem, offsetof(Ureg, ip), &amp;pc) &lt; 0) {
			rerrstr(buf, sizeof buf);
			if(strstr(buf, &#34;exited&#34;) || strstr(buf, &#34;No such process&#34;))
				return n;
			sysfatal(&#34;cannot read pc: %r&#34;);
		}
		pc--;
		if(put8(mem, offsetof(Ureg, ip), pc) &lt; 0)
			sysfatal(&#34;cannot write pc: %r&#34;);
		uncover(pc);
	}
}

void
main(int argc, char **argv)
{
	int n;

	ARGBEGIN{
	case &#39;g&#39;:
		substring = EARGF(usage());
		break;
	case &#39;l&#39;:
		longnames++;
		break;
	case &#39;n&#39;:
		minlines = atoi(EARGF(usage()));
		break;
	case &#39;s&#39;:
		doshowsrc = 1;
		break;
	case &#39;v&#39;:
		chatty++;
		break;
	default:
		usage();
	}ARGEND

	getwd(cwd, sizeof cwd);
	ncwd = strlen(cwd);

	if(argc == 0) {
		*--argv = &#34;6.out&#34;;
		argc++;
	}
	fd = open(argv[0], OREAD);
	if(fd &lt; 0)
		sysfatal(&#34;open %s: %r&#34;, argv[0]);
	if(crackhdr(fd, &amp;fhdr) &lt;= 0)
		sysfatal(&#34;crackhdr: %r&#34;);
	machbytype(fhdr.type);
	if(syminit(fd, &amp;fhdr) &lt;= 0)
		sysfatal(&#34;syminit: %r&#34;);
	text = loadmap(nil, fd, &amp;fhdr);
	if(text == nil)
		sysfatal(&#34;loadmap: %r&#34;);
	pid = startprocess(argv);
	mem = attachproc(pid, &amp;fhdr);
	if(mem == nil)
		sysfatal(&#34;attachproc: %r&#34;);
	breakpoints.cmp = rangecmp;
	cover();
	n = go();
	walktree(breakpoints.root);
	if(chatty)
		print(&#34;%d breakpoints\n&#34;, n);
	detachproc(mem);
	exits(0);
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

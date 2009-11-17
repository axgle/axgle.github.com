<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/mgc0.c</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/mgc0.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Garbage collector -- step 0.
//
// Stop the world, mark and sweep garbage collector.
// NOT INTENDED FOR PRODUCTION USE.
//
// A mark and sweep collector provides a way to exercise
// and test the memory allocator and the stack walking machinery
// without also needing to get reference counting
// exactly right.

#include &#34;runtime.h&#34;
#include &#34;malloc.h&#34;

enum {
	Debug = 0
};

extern byte data[];
extern byte etext[];
extern byte end[];

enum {
	PtrSize = sizeof(void*)
};

static void
scanblock(int32 depth, byte *b, int64 n)
{
	int32 off;
	void *obj;
	uintptr size;
	uint32 *ref;
	void **vp;
	int64 i;

	if(Debug)
		printf(&#34;%d scanblock %p %D\n&#34;, depth, b, n);
	off = (uint32)(uintptr)b &amp; (PtrSize-1);
	if(off) {
		b += PtrSize - off;
		n -= PtrSize - off;
	}

	vp = (void**)b;
	n /= PtrSize;
	for(i=0; i&lt;n; i++) {
		if(mlookup(vp[i], &amp;obj, &amp;size, &amp;ref)) {
			if(*ref == RefFree || *ref == RefStack)
				continue;
			if(*ref == RefNone) {
				if(Debug)
					printf(&#34;%d found at %p: &#34;, depth, &amp;vp[i]);
				*ref = RefSome;
				scanblock(depth+1, obj, size);
			}
		}
	}
}

static void
scanstack(G *gp)
{
	Stktop *stk;
	byte *sp;

	if(gp == g)
		sp = (byte*)&amp;gp;
	else
		sp = gp-&gt;sched.sp;
	stk = (Stktop*)gp-&gt;stackbase;
	while(stk) {
		scanblock(0, sp, (byte*)stk - sp);
		sp = stk-&gt;gobuf.sp;
		stk = (Stktop*)stk-&gt;stackbase;
	}
}

static void
mark(void)
{
	G *gp;

	// mark data+bss
	scanblock(0, data, end - data);

	// mark stacks
	for(gp=allg; gp!=nil; gp=gp-&gt;alllink) {
		switch(gp-&gt;status){
		default:
			printf(&#34;unexpected G.status %d\n&#34;, gp-&gt;status);
			throw(&#34;mark - bad status&#34;);
		case Gdead:
			break;
		case Grunning:
			if(gp != g)
				throw(&#34;mark - world not stopped&#34;);
			scanstack(gp);
			break;
		case Grunnable:
		case Gsyscall:
		case Gwaiting:
			scanstack(gp);
			break;
		}
	}
}

static void
sweepspan(MSpan *s)
{
	int32 i, n, npages, size;
	byte *p;

	if(s-&gt;state != MSpanInUse)
		return;

	p = (byte*)(s-&gt;start &lt;&lt; PageShift);
	if(s-&gt;sizeclass == 0) {
		// Large block.
		switch(s-&gt;gcref0) {
		default:
			throw(&#34;bad &#39;ref count&#39;&#34;);
		case RefFree:
		case RefManual:
		case RefStack:
			break;
		case RefNone:
			if(Debug)
				printf(&#34;free %D at %p\n&#34;, (uint64)s-&gt;npages&lt;&lt;PageShift, p);
			free(p);
			break;
		case RefSome:
//printf(&#34;gc-mem 1 %D\n&#34;, (uint64)s-&gt;npages&lt;&lt;PageShift);
			s-&gt;gcref0 = RefNone;	// set up for next mark phase
			break;
		}
		return;
	}

	// Chunk full of small blocks.
	// Must match computation in MCentral_Grow.
	size = class_to_size[s-&gt;sizeclass];
	npages = class_to_allocnpages[s-&gt;sizeclass];
	n = (npages &lt;&lt; PageShift) / (size + RefcountOverhead);
	for(i=0; i&lt;n; i++) {
		switch(s-&gt;gcref[i]) {
		default:
			throw(&#34;bad &#39;ref count&#39;&#34;);
		case RefFree:
		case RefManual:
		case RefStack:
			break;
		case RefNone:
			if(Debug)
				printf(&#34;free %d at %p\n&#34;, size, p+i*size);
			free(p + i*size);
			break;
		case RefSome:
			s-&gt;gcref[i] = RefNone;	// set up for next mark phase
			break;
		}
	}
//printf(&#34;gc-mem %d %d\n&#34;, s-&gt;ref, size);
}

static void
sweep(void)
{
	MSpan *s;

	// Sweep all the spans.
	for(s = mheap.allspans; s != nil; s = s-&gt;allnext)
		sweepspan(s);
}

// Semaphore, not Lock, so that the goroutine
// reschedules when there is contention rather
// than spinning.
static uint32 gcsema = 1;

// Initialized from $GOGC.  GOGC=off means no gc.
//
// Next gc is after we&#39;ve allocated an extra amount of
// memory proportional to the amount already in use.
// If gcpercent=100 and we&#39;re using 4M, we&#39;ll gc again
// when we get to 8M.  This keeps the gc cost in linear
// proportion to the allocation cost.  Adjusting gcpercent
// just changes the linear constant (and also the amount of
// extra memory used).
static int32 gcpercent = -2;

void
gc(int32 force)
{
	byte *p;

	// The gc is turned off (via enablegc) until
	// the bootstrap has completed.
	// Also, malloc gets called in the guts
	// of a number of libraries that might be
	// holding locks.  To avoid priority inversion
	// problems, don&#39;t bother trying to run gc
	// while holding a lock.  The next mallocgc
	// without a lock will do the gc instead.
	if(!mstats.enablegc || m-&gt;locks &gt; 0 || panicking)
		return;

	if(gcpercent == -2) {	// first time through
		p = getenv(&#34;GOGC&#34;);
		if(p == nil || p[0] == &#39;\0&#39;)
			gcpercent = 100;
		else if(strcmp(p, (byte*)&#34;off&#34;) == 0)
			gcpercent = -1;
		else
			gcpercent = atoi(p);
	}
	if(gcpercent &lt; 0)
		return;

//printf(&#34;gc...\n&#34;);
	semacquire(&amp;gcsema);
	m-&gt;gcing = 1;
	stoptheworld();
	if(mheap.Lock.key != 0)
		throw(&#34;mheap locked during gc&#34;);
	if(force || mstats.inuse_pages &gt;= mstats.next_gc) {
		mark();
		sweep();
		mstats.next_gc = mstats.inuse_pages+mstats.inuse_pages*gcpercent/100;
	}
	starttheworld();
	m-&gt;gcing = 0;
	semrelease(&amp;gcsema);
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

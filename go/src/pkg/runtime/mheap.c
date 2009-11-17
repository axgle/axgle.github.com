<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/mheap.c</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/mheap.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Page heap.
//
// See malloc.h for overview.
//
// When a MSpan is in the heap free list, state == MSpanFree
// and heapmap(s-&gt;start) == span, heapmap(s-&gt;start+s-&gt;npages-1) == span.
//
// When a MSpan is allocated, state == MSpanInUse
// and heapmap(i) == span for all s-&gt;start &lt;= i &lt; s-&gt;start+s-&gt;npages.

#include &#34;runtime.h&#34;
#include &#34;malloc.h&#34;

static MSpan *MHeap_AllocLocked(MHeap*, uintptr, int32);
static bool MHeap_Grow(MHeap*, uintptr);
static void MHeap_FreeLocked(MHeap*, MSpan*);
static MSpan *MHeap_AllocLarge(MHeap*, uintptr);
static MSpan *BestFit(MSpan*, uintptr, MSpan*);

static void
RecordSpan(void *vh, byte *p)
{
	MHeap *h;
	MSpan *s;

	h = vh;
	s = (MSpan*)p;
	s-&gt;allnext = h-&gt;allspans;
	h-&gt;allspans = s;
}

// Initialize the heap; fetch memory using alloc.
void
MHeap_Init(MHeap *h, void *(*alloc)(uintptr))
{
	uint32 i;

	FixAlloc_Init(&amp;h-&gt;spanalloc, sizeof(MSpan), alloc, RecordSpan, h);
	FixAlloc_Init(&amp;h-&gt;cachealloc, sizeof(MCache), alloc, nil, nil);
	MHeapMap_Init(&amp;h-&gt;map, alloc);
	// h-&gt;mapcache needs no init
	for(i=0; i&lt;nelem(h-&gt;free); i++)
		MSpanList_Init(&amp;h-&gt;free[i]);
	MSpanList_Init(&amp;h-&gt;large);
	for(i=0; i&lt;nelem(h-&gt;central); i++)
		MCentral_Init(&amp;h-&gt;central[i], i);
}

// Allocate a new span of npage pages from the heap
// and record its size class in the HeapMap and HeapMapCache.
MSpan*
MHeap_Alloc(MHeap *h, uintptr npage, int32 sizeclass)
{
	MSpan *s;

	lock(h);
	s = MHeap_AllocLocked(h, npage, sizeclass);
	if(s != nil)
		mstats.inuse_pages += npage;
	unlock(h);
	return s;
}

static MSpan*
MHeap_AllocLocked(MHeap *h, uintptr npage, int32 sizeclass)
{
	uintptr n;
	MSpan *s, *t;

	// Try in fixed-size lists up to max.
	for(n=npage; n &lt; nelem(h-&gt;free); n++) {
		if(!MSpanList_IsEmpty(&amp;h-&gt;free[n])) {
			s = h-&gt;free[n].next;
			goto HaveSpan;
		}
	}

	// Best fit in list of large spans.
	if((s = MHeap_AllocLarge(h, npage)) == nil) {
		if(!MHeap_Grow(h, npage))
			return nil;
		if((s = MHeap_AllocLarge(h, npage)) == nil)
			return nil;
	}

HaveSpan:
	// Mark span in use.
	if(s-&gt;state != MSpanFree)
		throw(&#34;MHeap_AllocLocked - MSpan not free&#34;);
	if(s-&gt;npages &lt; npage)
		throw(&#34;MHeap_AllocLocked - bad npages&#34;);
	MSpanList_Remove(s);
	s-&gt;state = MSpanInUse;

	if(s-&gt;npages &gt; npage) {
		// Trim extra and put it back in the heap.
		t = FixAlloc_Alloc(&amp;h-&gt;spanalloc);
		MSpan_Init(t, s-&gt;start + npage, s-&gt;npages - npage);
		s-&gt;npages = npage;
		MHeapMap_Set(&amp;h-&gt;map, t-&gt;start - 1, s);
		MHeapMap_Set(&amp;h-&gt;map, t-&gt;start, t);
		MHeapMap_Set(&amp;h-&gt;map, t-&gt;start + t-&gt;npages - 1, t);
		t-&gt;state = MSpanInUse;
		MHeap_FreeLocked(h, t);
	}

	// If span is being used for small objects, cache size class.
	// No matter what, cache span info, because gc needs to be
	// able to map interior pointer to containing span.
	s-&gt;sizeclass = sizeclass;
	for(n=0; n&lt;npage; n++)
		MHeapMap_Set(&amp;h-&gt;map, s-&gt;start+n, s);
	if(sizeclass == 0) {
		uintptr tmp;

		// If there are entries for this span, invalidate them,
		// but don&#39;t blow out cache entries about other spans.
		for(n=0; n&lt;npage; n++)
			if(MHeapMapCache_GET(&amp;h-&gt;mapcache, s-&gt;start+n, tmp) != 0)
				MHeapMapCache_SET(&amp;h-&gt;mapcache, s-&gt;start+n, 0);
	} else {
		// Save cache entries for this span.
		// If there&#39;s a size class, there aren&#39;t that many pages.
		for(n=0; n&lt;npage; n++)
			MHeapMapCache_SET(&amp;h-&gt;mapcache, s-&gt;start+n, sizeclass);
	}

	return s;
}

// Allocate a span of exactly npage pages from the list of large spans.
static MSpan*
MHeap_AllocLarge(MHeap *h, uintptr npage)
{
	return BestFit(&amp;h-&gt;large, npage, nil);
}

// Search list for smallest span with &gt;= npage pages.
// If there are multiple smallest spans, take the one
// with the earliest starting address.
static MSpan*
BestFit(MSpan *list, uintptr npage, MSpan *best)
{
	MSpan *s;

	for(s=list-&gt;next; s != list; s=s-&gt;next) {
		if(s-&gt;npages &lt; npage)
			continue;
		if(best == nil
		|| s-&gt;npages &lt; best-&gt;npages
		|| (s-&gt;npages == best-&gt;npages &amp;&amp; s-&gt;start &lt; best-&gt;start))
			best = s;
	}
	return best;
}

// Try to add at least npage pages of memory to the heap,
// returning whether it worked.
static bool
MHeap_Grow(MHeap *h, uintptr npage)
{
	uintptr ask;
	void *v;
	MSpan *s;

	// Ask for a big chunk, to reduce the number of mappings
	// the operating system needs to track; also amortizes
	// the overhead of an operating system mapping.
	// For Native Client, allocate a multiple of 64kB (16 pages).
	npage = (npage+15)&amp;~15;
	ask = npage&lt;&lt;PageShift;
	if(ask &lt; HeapAllocChunk)
		ask = HeapAllocChunk;

	v = SysAlloc(ask);
	if(v == nil) {
		if(ask &gt; (npage&lt;&lt;PageShift)) {
			ask = npage&lt;&lt;PageShift;
			v = SysAlloc(ask);
		}
		if(v == nil)
			return false;
	}

	// NOTE(rsc): In tcmalloc, if we&#39;ve accumulated enough
	// system allocations, the heap map gets entirely allocated
	// in 32-bit mode.  (In 64-bit mode that&#39;s not practical.)

	if(!MHeapMap_Preallocate(&amp;h-&gt;map, ((uintptr)v&gt;&gt;PageShift) - 1, (ask&gt;&gt;PageShift) + 2)) {
		SysFree(v, ask);
		return false;
	}

	// Create a fake &#34;in use&#34; span and free it, so that the
	// right coalescing happens.
	s = FixAlloc_Alloc(&amp;h-&gt;spanalloc);
	MSpan_Init(s, (uintptr)v&gt;&gt;PageShift, ask&gt;&gt;PageShift);
	MHeapMap_Set(&amp;h-&gt;map, s-&gt;start, s);
	MHeapMap_Set(&amp;h-&gt;map, s-&gt;start + s-&gt;npages - 1, s);
	s-&gt;state = MSpanInUse;
	MHeap_FreeLocked(h, s);
	return true;
}

// Look up the span at the given page number.
// Page number is guaranteed to be in map
// and is guaranteed to be start or end of span.
MSpan*
MHeap_Lookup(MHeap *h, PageID p)
{
	return MHeapMap_Get(&amp;h-&gt;map, p);
}

// Look up the span at the given page number.
// Page number is *not* guaranteed to be in map
// and may be anywhere in the span.
// Map entries for the middle of a span are only
// valid for allocated spans.  Free spans may have
// other garbage in their middles, so we have to
// check for that.
MSpan*
MHeap_LookupMaybe(MHeap *h, PageID p)
{
	MSpan *s;

	s = MHeapMap_GetMaybe(&amp;h-&gt;map, p);
	if(s == nil || p &lt; s-&gt;start || p - s-&gt;start &gt;= s-&gt;npages)
		return nil;
	if(s-&gt;state != MSpanInUse)
		return nil;
	return s;
}

// Free the span back into the heap.
void
MHeap_Free(MHeap *h, MSpan *s)
{
	lock(h);
	mstats.inuse_pages -= s-&gt;npages;
	MHeap_FreeLocked(h, s);
	unlock(h);
}

static void
MHeap_FreeLocked(MHeap *h, MSpan *s)
{
	MSpan *t;

	if(s-&gt;state != MSpanInUse || s-&gt;ref != 0) {
		printf(&#34;MHeap_FreeLocked - span %p ptr %p state %d ref %d\n&#34;, s, s-&gt;start&lt;&lt;PageShift, s-&gt;state, s-&gt;ref);
		throw(&#34;MHeap_FreeLocked - invalid free&#34;);
	}
	s-&gt;state = MSpanFree;
	MSpanList_Remove(s);

	// Coalesce with earlier, later spans.
	if((t = MHeapMap_Get(&amp;h-&gt;map, s-&gt;start - 1)) != nil &amp;&amp; t-&gt;state != MSpanInUse) {
		s-&gt;start = t-&gt;start;
		s-&gt;npages += t-&gt;npages;
		MHeapMap_Set(&amp;h-&gt;map, s-&gt;start, s);
		MSpanList_Remove(t);
		t-&gt;state = MSpanDead;
		FixAlloc_Free(&amp;h-&gt;spanalloc, t);
	}
	if((t = MHeapMap_Get(&amp;h-&gt;map, s-&gt;start + s-&gt;npages)) != nil &amp;&amp; t-&gt;state != MSpanInUse) {
		s-&gt;npages += t-&gt;npages;
		MHeapMap_Set(&amp;h-&gt;map, s-&gt;start + s-&gt;npages - 1, s);
		MSpanList_Remove(t);
		t-&gt;state = MSpanDead;
		FixAlloc_Free(&amp;h-&gt;spanalloc, t);
	}

	// Insert s into appropriate list.
	if(s-&gt;npages &lt; nelem(h-&gt;free))
		MSpanList_Insert(&amp;h-&gt;free[s-&gt;npages], s);
	else
		MSpanList_Insert(&amp;h-&gt;large, s);

	// TODO(rsc): IncrementalScavenge() to return memory to OS.
}

// Initialize a new span with the given start and npages.
void
MSpan_Init(MSpan *span, PageID start, uintptr npages)
{
	span-&gt;next = nil;
	span-&gt;prev = nil;
	span-&gt;start = start;
	span-&gt;npages = npages;
	span-&gt;freelist = nil;
	span-&gt;ref = 0;
	span-&gt;sizeclass = 0;
	span-&gt;state = 0;
}

// Initialize an empty doubly-linked list.
void
MSpanList_Init(MSpan *list)
{
	list-&gt;state = MSpanListHead;
	list-&gt;next = list;
	list-&gt;prev = list;
}

void
MSpanList_Remove(MSpan *span)
{
	if(span-&gt;prev == nil &amp;&amp; span-&gt;next == nil)
		return;
	span-&gt;prev-&gt;next = span-&gt;next;
	span-&gt;next-&gt;prev = span-&gt;prev;
	span-&gt;prev = nil;
	span-&gt;next = nil;
}

bool
MSpanList_IsEmpty(MSpan *list)
{
	return list-&gt;next == list;
}

void
MSpanList_Insert(MSpan *list, MSpan *span)
{
	if(span-&gt;next != nil || span-&gt;prev != nil)
		throw(&#34;MSpanList_Insert&#34;);
	span-&gt;next = list-&gt;next;
	span-&gt;prev = list;
	span-&gt;next-&gt;prev = span;
	span-&gt;prev-&gt;next = span;
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

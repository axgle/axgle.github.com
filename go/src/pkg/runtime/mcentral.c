<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/mcentral.c</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/mcentral.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Central free lists.
//
// See malloc.h for an overview.
//
// The MCentral doesn&#39;t actually contain the list of free objects; the MSpan does.
// Each MCentral is two lists of MSpans: those with free objects (c-&gt;nonempty)
// and those that are completely allocated (c-&gt;empty).
//
// TODO(rsc): tcmalloc uses a &#34;transfer cache&#34; to split the list
// into sections of class_to_transfercount[sizeclass] objects
// so that it is faster to move those lists between MCaches and MCentrals.

#include &#34;runtime.h&#34;
#include &#34;malloc.h&#34;

static bool MCentral_Grow(MCentral *c);
static void* MCentral_Alloc(MCentral *c);
static void MCentral_Free(MCentral *c, void *v);

// Initialize a single central free list.
void
MCentral_Init(MCentral *c, int32 sizeclass)
{
	c-&gt;sizeclass = sizeclass;
	MSpanList_Init(&amp;c-&gt;nonempty);
	MSpanList_Init(&amp;c-&gt;empty);
}

// Allocate up to n objects from the central free list.
// Return the number of objects allocated.
// The objects are linked together by their first words.
// On return, *pstart points at the first object and *pend at the last.
int32
MCentral_AllocList(MCentral *c, int32 n, MLink **pfirst)
{
	MLink *first, *last, *v;
	int32 i;


	lock(c);
	// Replenish central list if empty.
	if(MSpanList_IsEmpty(&amp;c-&gt;nonempty)) {
		if(!MCentral_Grow(c)) {
			unlock(c);
			*pfirst = nil;
			return 0;
		}
	}

	// Copy from list, up to n.
	// First one is guaranteed to work, because we just grew the list.
	first = MCentral_Alloc(c);
	last = first;
	for(i=1; i&lt;n &amp;&amp; (v = MCentral_Alloc(c)) != nil; i++) {
		last-&gt;next = v;
		last = v;
	}
	last-&gt;next = nil;
	c-&gt;nfree -= i;

	unlock(c);
	*pfirst = first;
	return i;
}

// Helper: allocate one object from the central free list.
static void*
MCentral_Alloc(MCentral *c)
{
	MSpan *s;
	MLink *v;

	if(MSpanList_IsEmpty(&amp;c-&gt;nonempty))
		return nil;
	s = c-&gt;nonempty.next;
	s-&gt;ref++;
	v = s-&gt;freelist;
	s-&gt;freelist = v-&gt;next;
	if(s-&gt;freelist == nil) {
		MSpanList_Remove(s);
		MSpanList_Insert(&amp;c-&gt;empty, s);
	}
	return v;
}

// Free n objects back into the central free list.
// Return the number of objects allocated.
// The objects are linked together by their first words.
// On return, *pstart points at the first object and *pend at the last.
void
MCentral_FreeList(MCentral *c, int32 n, MLink *start)
{
	MLink *v, *next;

	// Assume next == nil marks end of list.
	// n and end would be useful if we implemented
	// the transfer cache optimization in the TODO above.
	USED(n);

	lock(c);
	for(v=start; v; v=next) {
		next = v-&gt;next;
		MCentral_Free(c, v);
	}
	unlock(c);
}

// Helper: free one object back into the central free list.
static void
MCentral_Free(MCentral *c, void *v)
{
	MSpan *s;
	PageID page;
	MLink *p, *next;

	// Find span for v.
	page = (uintptr)v &gt;&gt; PageShift;
	s = MHeap_Lookup(&amp;mheap, page);
	if(s == nil || s-&gt;ref == 0)
		throw(&#34;invalid free&#34;);

	// Move to nonempty if necessary.
	if(s-&gt;freelist == nil) {
		MSpanList_Remove(s);
		MSpanList_Insert(&amp;c-&gt;nonempty, s);
	}

	// Add v back to s&#39;s free list.
	p = v;
	p-&gt;next = s-&gt;freelist;
	s-&gt;freelist = p;
	c-&gt;nfree++;

	// If s is completely freed, return it to the heap.
	if(--s-&gt;ref == 0) {
		MSpanList_Remove(s);
		// Freed blocks are zeroed except for the link pointer.
		// Zero the link pointers so that the page is all zero.
		for(p=s-&gt;freelist; p; p=next) {
			next = p-&gt;next;
			p-&gt;next = nil;
		}
		s-&gt;freelist = nil;
		c-&gt;nfree -= (s-&gt;npages &lt;&lt; PageShift) / class_to_size[c-&gt;sizeclass];
		unlock(c);
		MHeap_Free(&amp;mheap, s);
		lock(c);
	}
}

// Fetch a new span from the heap and
// carve into objects for the free list.
static bool
MCentral_Grow(MCentral *c)
{
	int32 i, n, npages, size;
	MLink **tailp, *v;
	byte *p;
	MSpan *s;

	unlock(c);
	npages = class_to_allocnpages[c-&gt;sizeclass];
	s = MHeap_Alloc(&amp;mheap, npages, c-&gt;sizeclass);
	if(s == nil) {
		// TODO(rsc): Log out of memory
		lock(c);
		return false;
	}

	// Carve span into sequence of blocks.
	tailp = &amp;s-&gt;freelist;
	p = (byte*)(s-&gt;start &lt;&lt; PageShift);
	size = class_to_size[c-&gt;sizeclass];
	n = (npages &lt;&lt; PageShift) / (size + RefcountOverhead);
	s-&gt;gcref = (uint32*)(p + size*n);
	for(i=0; i&lt;n; i++) {
		v = (MLink*)p;
		*tailp = v;
		tailp = &amp;v-&gt;next;
		p += size;
	}
	*tailp = nil;

	lock(c);
	c-&gt;nfree += n;
	MSpanList_Insert(&amp;c-&gt;nonempty, s);
	return true;
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

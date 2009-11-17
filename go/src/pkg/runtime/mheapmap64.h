<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/mheapmap64.h</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/mheapmap64.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Free(v) must be able to determine the MSpan containing v.
// The MHeapMap is a 3-level radix tree mapping page numbers to MSpans.
//
// NOTE(rsc): On a 32-bit platform (= 20-bit page numbers),
// we can swap in a 2-level radix tree.
//
// NOTE(rsc): We use a 3-level tree because tcmalloc does, but
// having only three levels requires approximately 1 MB per node
// in the tree, making the minimum map footprint 3 MB.
// Using a 4-level tree would cut the minimum footprint to 256 kB.
// On the other hand, it&#39;s just virtual address space: most of
// the memory is never going to be touched, thus never paged in.

typedef struct MHeapMapNode2 MHeapMapNode2;
typedef struct MHeapMapNode3 MHeapMapNode3;

enum
{
	// 64 bit address - 12 bit page size = 52 bits to map
	MHeapMap_Level1Bits = 18,
	MHeapMap_Level2Bits = 18,
	MHeapMap_Level3Bits = 16,

	MHeapMap_TotalBits =
		MHeapMap_Level1Bits +
		MHeapMap_Level2Bits +
		MHeapMap_Level3Bits,

	MHeapMap_Level1Mask = (1&lt;&lt;MHeapMap_Level1Bits) - 1,
	MHeapMap_Level2Mask = (1&lt;&lt;MHeapMap_Level2Bits) - 1,
	MHeapMap_Level3Mask = (1&lt;&lt;MHeapMap_Level3Bits) - 1,
};

struct MHeapMap
{
	void *(*allocator)(uintptr);
	MHeapMapNode2 *p[1&lt;&lt;MHeapMap_Level1Bits];
};

struct MHeapMapNode2
{
	MHeapMapNode3 *p[1&lt;&lt;MHeapMap_Level2Bits];
};

struct MHeapMapNode3
{
	MSpan *s[1&lt;&lt;MHeapMap_Level3Bits];
};

void	MHeapMap_Init(MHeapMap *m, void *(*allocator)(uintptr));
bool	MHeapMap_Preallocate(MHeapMap *m, PageID k, uintptr npages);
MSpan*	MHeapMap_Get(MHeapMap *m, PageID k);
MSpan*	MHeapMap_GetMaybe(MHeapMap *m, PageID k);
void	MHeapMap_Set(MHeapMap *m, PageID k, MSpan *v);


// Much of the time, free(v) needs to know only the size class for v,
// not which span it came from.  The MHeapMap finds the size class
// by looking up the span.
//
// An MHeapMapCache is a simple direct-mapped cache translating
// page numbers to size classes.  It avoids the expensive MHeapMap
// lookup for hot pages.
//
// The cache entries are 64 bits, with the page number in the low part
// and the value at the top.
//
// NOTE(rsc): On a machine with 32-bit addresses (= 20-bit page numbers),
// we can use a 16-bit cache entry by not storing the redundant 12 bits
// of the key that are used as the entry index.  Here in 64-bit land,
// that trick won&#39;t work unless the hash table has 2^28 entries.
enum
{
	MHeapMapCache_HashBits = 12
};

struct MHeapMapCache
{
	uintptr array[1&lt;&lt;MHeapMapCache_HashBits];
};

// All macros for speed (sorry).
#define HMASK	((1&lt;&lt;MHeapMapCache_HashBits)-1)
#define KBITS	MHeapMap_TotalBits
#define KMASK	((1LL&lt;&lt;KBITS)-1)

#define MHeapMapCache_SET(cache, key, value) \
	((cache)-&gt;array[(key) &amp; HMASK] = (key) | ((uintptr)(value) &lt;&lt; KBITS))

#define MHeapMapCache_GET(cache, key, tmp) \
	(tmp = (cache)-&gt;array[(key) &amp; HMASK], \
	 (tmp &amp; KMASK) == (key) ? (tmp &gt;&gt; KBITS) : 0)
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

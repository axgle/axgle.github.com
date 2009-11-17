<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/mcache.c</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/mcache.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Per-thread (in Go, per-M) malloc cache for small objects.
//
// See malloc.h for an overview.

#include &#34;runtime.h&#34;
#include &#34;malloc.h&#34;

void*
MCache_Alloc(MCache *c, int32 sizeclass, uintptr size)
{
	MCacheList *l;
	MLink *first, *v;
	int32 n;

	// Allocate from list.
	l = &amp;c-&gt;list[sizeclass];
	if(l-&gt;list == nil) {
		// Replenish using central lists.
		n = MCentral_AllocList(&amp;mheap.central[sizeclass],
			class_to_transfercount[sizeclass], &amp;first);
		l-&gt;list = first;
		l-&gt;nlist = n;
		c-&gt;size += n*size;
	}
	v = l-&gt;list;
	l-&gt;list = v-&gt;next;
	l-&gt;nlist--;
	if(l-&gt;nlist &lt; l-&gt;nlistmin)
		l-&gt;nlistmin = l-&gt;nlist;
	c-&gt;size -= size;

	// v is zeroed except for the link pointer
	// that we used above; zero that.
	v-&gt;next = nil;
	return v;
}

// Take n elements off l and return them to the central free list.
static void
ReleaseN(MCache *c, MCacheList *l, int32 n, int32 sizeclass)
{
	MLink *first, **lp;
	int32 i;

	// Cut off first n elements.
	first = l-&gt;list;
	lp = &amp;l-&gt;list;
	for(i=0; i&lt;n; i++)
		lp = &amp;(*lp)-&gt;next;
	l-&gt;list = *lp;
	*lp = nil;
	l-&gt;nlist -= n;
	if(l-&gt;nlist &lt; l-&gt;nlistmin)
		l-&gt;nlistmin = l-&gt;nlist;
	c-&gt;size -= n*class_to_size[sizeclass];

	// Return them to central free list.
	MCentral_FreeList(&amp;mheap.central[sizeclass], n, first);
}

void
MCache_Free(MCache *c, void *v, int32 sizeclass, uintptr size)
{
	int32 i, n;
	MCacheList *l;
	MLink *p;

	// Put back on list.
	l = &amp;c-&gt;list[sizeclass];
	p = v;
	p-&gt;next = l-&gt;list;
	l-&gt;list = p;
	l-&gt;nlist++;
	c-&gt;size += size;

	if(l-&gt;nlist &gt;= MaxMCacheListLen) {
		// Release a chunk back.
		ReleaseN(c, l, class_to_transfercount[sizeclass], sizeclass);
	}

	if(c-&gt;size &gt;= MaxMCacheSize) {
		// Scavenge.
		for(i=0; i&lt;NumSizeClasses; i++) {
			l = &amp;c-&gt;list[i];
			n = l-&gt;nlistmin;

			// n is the minimum number of elements we&#39;ve seen on
			// the list since the last scavenge.  If n &gt; 0, it means that
			// we could have gotten by with n fewer elements
			// without needing to consult the central free list.
			// Move toward that situation by releasing n/2 of them.
			if(n &gt; 0) {
				if(n &gt; 1)
					n /= 2;
				ReleaseN(c, l, n, i);
			}
			l-&gt;nlistmin = l-&gt;nlist;
		}
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

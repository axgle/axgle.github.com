<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/mheapmap64.c</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/mheapmap64.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Heap map, 64-bit version 
// See malloc.h and mheap.c for overview.

#include &#34;runtime.h&#34;
#include &#34;malloc.h&#34;

// 3-level radix tree mapping page ids to Span*.
void
MHeapMap_Init(MHeapMap *m, void *(*allocator)(size_t))
{
	m-&gt;allocator = allocator;
}

MSpan*
MHeapMap_Get(MHeapMap *m, PageID k)
{
	int32 i1, i2, i3;

	i3 = k &amp; MHeapMap_Level3Mask;
	k &gt;&gt;= MHeapMap_Level3Bits;
	i2 = k &amp; MHeapMap_Level2Mask;
	k &gt;&gt;= MHeapMap_Level2Bits;
	i1 = k &amp; MHeapMap_Level1Mask;
	k &gt;&gt;= MHeapMap_Level1Bits;
	if(k != 0)
		throw(&#34;MHeapMap_Get&#34;);

	return m-&gt;p[i1]-&gt;p[i2]-&gt;s[i3];
}

MSpan*
MHeapMap_GetMaybe(MHeapMap *m, PageID k)
{
	int32 i1, i2, i3;
	MHeapMapNode2 *p2;
	MHeapMapNode3 *p3;

	i3 = k &amp; MHeapMap_Level3Mask;
	k &gt;&gt;= MHeapMap_Level3Bits;
	i2 = k &amp; MHeapMap_Level2Mask;
	k &gt;&gt;= MHeapMap_Level2Bits;
	i1 = k &amp; MHeapMap_Level1Mask;
	k &gt;&gt;= MHeapMap_Level1Bits;
	if(k != 0)
		throw(&#34;MHeapMap_Get&#34;);

	p2 = m-&gt;p[i1];
	if(p2 == nil)
		return nil;
	p3 = p2-&gt;p[i2];
	if(p3 == nil)
		return nil;
	return p3-&gt;s[i3];
}

void
MHeapMap_Set(MHeapMap *m, PageID k, MSpan *s)
{
	int32 i1, i2, i3;

	i3 = k &amp; MHeapMap_Level3Mask;
	k &gt;&gt;= MHeapMap_Level3Bits;
	i2 = k &amp; MHeapMap_Level2Mask;
	k &gt;&gt;= MHeapMap_Level2Bits;
	i1 = k &amp; MHeapMap_Level1Mask;
	k &gt;&gt;= MHeapMap_Level1Bits;
	if(k != 0)
		throw(&#34;MHeapMap_Set&#34;);

	m-&gt;p[i1]-&gt;p[i2]-&gt;s[i3] = s;
}

// Allocate the storage required for entries [k, k+1, ..., k+len-1]
// so that Get and Set calls need not check for nil pointers.
bool
MHeapMap_Preallocate(MHeapMap *m, PageID k, uintptr len)
{
	uintptr end;
	int32 i1, i2;
	MHeapMapNode2 *p2;
	MHeapMapNode3 *p3;

	end = k+len;
	while(k &lt; end) {
		if((k &gt;&gt; MHeapMap_TotalBits) != 0)
			return false;
		i2 = (k &gt;&gt; MHeapMap_Level3Bits) &amp; MHeapMap_Level2Mask;
		i1 = (k &gt;&gt; (MHeapMap_Level3Bits + MHeapMap_Level2Bits)) &amp; MHeapMap_Level1Mask;

		// first-level pointer
		if((p2 = m-&gt;p[i1]) == nil) {
			p2 = m-&gt;allocator(sizeof *p2);
			if(p2 == nil)
				return false;
			runtime_memclr((byte*)p2, sizeof *p2);
			m-&gt;p[i1] = p2;
		}

		// second-level pointer
		if(p2-&gt;p[i2] == nil) {
			p3 = m-&gt;allocator(sizeof *p3);
			if(p3 == nil)
				return false;
			runtime_memclr((byte*)p3, sizeof *p3);
			p2-&gt;p[i2] = p3;
		}

		// advance key past this leaf node
		k = ((k &gt;&gt; MHeapMap_Level3Bits) + 1) &lt;&lt; MHeapMap_Level3Bits;
	}
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

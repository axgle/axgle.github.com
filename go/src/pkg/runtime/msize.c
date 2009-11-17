<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/msize.c</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/msize.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Malloc small size classes.
//
// See malloc.h for overview.
//
// The size classes are chosen so that rounding an allocation
// request up to the next size class wastes at most 12.5% (1.125x).
//
// Each size class has its own page count that gets allocated
// and chopped up when new objects of the size class are needed.
// That page count is chosen so that chopping up the run of
// pages into objects of the given size wastes at most 12.5% (1.125x)
// of the memory.  It is not necessary that the cutoff here be
// the same as above.
//
// The two sources of waste multiply, so the worst possible case
// for the above constraints would be that allocations of some
// size might have a 26.6% (1.266x) overhead.
// In practice, only one of the wastes comes into play for a
// given size (sizes &lt; 512 waste mainly on the round-up,
// sizes &gt; 512 waste mainly on the page chopping).
//
// TODO(rsc): Compute max waste for any given size.

#include &#34;runtime.h&#34;
#include &#34;malloc.h&#34;

int32 class_to_size[NumSizeClasses];
int32 class_to_allocnpages[NumSizeClasses];
int32 class_to_transfercount[NumSizeClasses];

// The SizeToClass lookup is implemented using two arrays,
// one mapping sizes &lt;= 1024 to their class and one mapping
// sizes &gt;= 1024 and &lt;= MaxSmallSize to their class.
// All objects are 8-aligned, so the first array is indexed by
// the size divided by 8 (rounded up).  Objects &gt;= 1024 bytes
// are 128-aligned, so the second array is indexed by the
// size divided by 128 (rounded up).  The arrays are filled in
// by InitSizes.

static int32 size_to_class8[1024/8 + 1];
static int32 size_to_class128[(MaxSmallSize-1024)/128 + 1];

int32
SizeToClass(int32 size)
{
	if(size &gt; MaxSmallSize)
		throw(&#34;SizeToClass - invalid size&#34;);
	if(size &gt; 1024-8)
		return size_to_class128[(size-1024+127) &gt;&gt; 7];
	return size_to_class8[(size+7)&gt;&gt;3];
}

void
InitSizes(void)
{
	int32 align, sizeclass, size, osize, nextsize, n;
	uint32 i;
	uintptr allocsize, npages;

	// Initialize the class_to_size table (and choose class sizes in the process).
	class_to_size[0] = 0;
	sizeclass = 1;	// 0 means no class
	align = 8;
	for(size = align; size &lt;= MaxSmallSize; size += align) {
		if((size&amp;(size-1)) == 0) {	// bump alignment once in a while
			if(size &gt;= 2048)
				align = 256;
			else if(size &gt;= 128)
				align = size / 8;
			else if(size &gt;= 16)
				align = 16;	// required for x86 SSE instructions, if we want to use them
		}
		if((align&amp;(align-1)) != 0)
			throw(&#34;InitSizes - bug&#34;);

		// Make the allocnpages big enough that
		// the leftover is less than 1/8 of the total,
		// so wasted space is at most 12.5%.
		allocsize = PageSize;
		osize = size + RefcountOverhead;
		while(allocsize%osize &gt; (allocsize/8))
			allocsize += PageSize;
		npages = allocsize &gt;&gt; PageShift;

		// If the previous sizeclass chose the same
		// allocation size and fit the same number of
		// objects into the page, we might as well
		// use just this size instead of having two
		// different sizes.
		if(sizeclass &gt; 1
		&amp;&amp; npages == class_to_allocnpages[sizeclass-1]
		&amp;&amp; allocsize/osize == allocsize/(class_to_size[sizeclass-1]+RefcountOverhead)) {
			class_to_size[sizeclass-1] = size;
			continue;
		}

		class_to_allocnpages[sizeclass] = npages;
		class_to_size[sizeclass] = size;
		sizeclass++;
	}
	if(sizeclass != NumSizeClasses) {
		printf(&#34;sizeclass=%d NumSizeClasses=%d\n&#34;, sizeclass, NumSizeClasses);
		throw(&#34;InitSizes - bad NumSizeClasses&#34;);
	}

	// Initialize the size_to_class tables.
	nextsize = 0;
	for (sizeclass = 1; sizeclass &lt; NumSizeClasses; sizeclass++) {
		for(; nextsize &lt; 1024 &amp;&amp; nextsize &lt;= class_to_size[sizeclass]; nextsize+=8)
			size_to_class8[nextsize/8] = sizeclass;
		if(nextsize &gt;= 1024)
			for(; nextsize &lt;= class_to_size[sizeclass]; nextsize += 128)
				size_to_class128[(nextsize-1024)/128] = sizeclass;
	}

	// Double-check SizeToClass.
	if(0) {
		for(n=0; n &lt; MaxSmallSize; n++) {
			sizeclass = SizeToClass(n);
			if(sizeclass &lt; 1 || sizeclass &gt;= NumSizeClasses || class_to_size[sizeclass] &lt; n) {
				printf(&#34;size=%d sizeclass=%d class_to_size=%d\n&#34;, n, sizeclass, class_to_size[sizeclass]);
				printf(&#34;incorrect SizeToClass&#34;);
				goto dump;
			}
			if(sizeclass &gt; 1 &amp;&amp; class_to_size[sizeclass-1] &gt;= n) {
				printf(&#34;size=%d sizeclass=%d class_to_size=%d\n&#34;, n, sizeclass, class_to_size[sizeclass]);
				printf(&#34;SizeToClass too big&#34;);
				goto dump;
			}
		}
	}

	// Initialize the class_to_transfercount table.
	for(sizeclass = 1; sizeclass &lt; NumSizeClasses; sizeclass++) {
		n = 64*1024 / class_to_size[sizeclass];
		if(n &lt; 2)
			n = 2;
		if(n &gt; 32)
			n = 32;
		class_to_transfercount[sizeclass] = n;
	}
	return;

dump:
	if(1){
		printf(&#34;NumSizeClasses=%d\n&#34;, NumSizeClasses);
		printf(&#34;class_to_size:&#34;);
		for(sizeclass=0; sizeclass&lt;NumSizeClasses; sizeclass++)
			printf(&#34; %d&#34;, class_to_size[sizeclass]);
		printf(&#34;\n\n&#34;);
		printf(&#34;size_to_class8:&#34;);
		for(i=0; i&lt;nelem(size_to_class8); i++)
			printf(&#34; %d=&gt;%d(%d)\n&#34;, i*8, size_to_class8[i], class_to_size[size_to_class8[i]]);
		printf(&#34;\n&#34;);
		printf(&#34;size_to_class128:&#34;);
		for(i=0; i&lt;nelem(size_to_class128); i++)
			printf(&#34; %d=&gt;%d(%d)\n&#34;, i*128, size_to_class128[i], class_to_size[size_to_class128[i]]);
		printf(&#34;\n&#34;);
	}
	throw(&#34;InitSizes failed&#34;);
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

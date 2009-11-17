<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/slice.c</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/slice.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;

static	int32	debug	= 0;

// makeslice(nel int, cap int, width int) (ary []any);
void
runtime·makeslice(uint32 nel, uint32 cap, uint32 width, Slice ret)
{
	uint64 size;

	if(cap &lt; nel)
		cap = nel;
	size = cap*width;

	ret.len = nel;
	ret.cap = cap;
	ret.array = mal(size);

	FLUSH(&amp;ret);

	if(debug) {
		prints(&#34;makeslice: nel=&#34;);
		runtime·printint(nel);
		prints(&#34;; cap=&#34;);
		runtime·printint(cap);
		prints(&#34;; width=&#34;);
		runtime·printint(width);
		prints(&#34;; ret=&#34;);
		runtime·printslice(ret);
		prints(&#34;\n&#34;);
	}
}

static void
throwslice(uint32 lb, uint32 hb, uint32 n)
{
	prints(&#34;slice[&#34;);
	runtime·printint(lb);
	prints(&#34;:&#34;);
	runtime·printint(hb);
	prints(&#34;] of [&#34;);
	runtime·printint(n);
	prints(&#34;] array\n&#34;);
	throw(&#34;array slice&#34;);
}

// sliceslice(old []any, lb int, hb int, width int) (ary []any);
void
runtime·sliceslice(Slice old, uint32 lb, uint32 hb, uint32 width, Slice ret)
{

	if(hb &gt; old.cap || lb &gt; hb) {
		if(debug) {
			prints(&#34;runtime·sliceslice: old=&#34;);
			runtime·printslice(old);
			prints(&#34;; lb=&#34;);
			runtime·printint(lb);
			prints(&#34;; hb=&#34;);
			runtime·printint(hb);
			prints(&#34;; width=&#34;);
			runtime·printint(width);
			prints(&#34;\n&#34;);

			prints(&#34;oldarray: nel=&#34;);
			runtime·printint(old.len);
			prints(&#34;; cap=&#34;);
			runtime·printint(old.cap);
			prints(&#34;\n&#34;);
		}
		throwslice(lb, hb, old.cap);
	}

	// new array is inside old array
	ret.len = hb-lb;
	ret.cap = old.cap - lb;
	ret.array = old.array + lb*width;

	FLUSH(&amp;ret);

	if(debug) {
		prints(&#34;runtime·sliceslice: old=&#34;);
		runtime·printslice(old);
		prints(&#34;; lb=&#34;);
		runtime·printint(lb);
		prints(&#34;; hb=&#34;);
		runtime·printint(hb);
		prints(&#34;; width=&#34;);
		runtime·printint(width);
		prints(&#34;; ret=&#34;);
		runtime·printslice(ret);
		prints(&#34;\n&#34;);
	}
}

// slicearray(old *any, nel int, lb int, hb int, width int) (ary []any);
void
runtime·slicearray(byte* old, uint32 nel, uint32 lb, uint32 hb, uint32 width, Slice ret)
{
	if(nel &gt; 0 &amp;&amp; old == nil) {
		// crash if old == nil.
		// could give a better message
		// but this is consistent with all the in-line checks
		// that the compiler inserts for other uses.
		*old = 0;
	}

	if(hb &gt; nel || lb &gt; hb) {
		if(debug) {
			prints(&#34;runtime·slicearray: old=&#34;);
			runtime·printpointer(old);
			prints(&#34;; nel=&#34;);
			runtime·printint(nel);
			prints(&#34;; lb=&#34;);
			runtime·printint(lb);
			prints(&#34;; hb=&#34;);
			runtime·printint(hb);
			prints(&#34;; width=&#34;);
			runtime·printint(width);
			prints(&#34;\n&#34;);
		}
		throwslice(lb, hb, nel);
	}

	// new array is inside old array
	ret.len = hb-lb;
	ret.cap = nel-lb;
	ret.array = old + lb*width;

	FLUSH(&amp;ret);

	if(debug) {
		prints(&#34;runtime·slicearray: old=&#34;);
		runtime·printpointer(old);
		prints(&#34;; nel=&#34;);
		runtime·printint(nel);
		prints(&#34;; lb=&#34;);
		runtime·printint(lb);
		prints(&#34;; hb=&#34;);
		runtime·printint(hb);
		prints(&#34;; width=&#34;);
		runtime·printint(width);
		prints(&#34;; ret=&#34;);
		runtime·printslice(ret);
		prints(&#34;\n&#34;);
	}
}

// arraytoslice(old *any, nel int) (ary []any)
void
runtime·arraytoslice(byte* old, uint32 nel, Slice ret)
{
	if(nel &gt; 0 &amp;&amp; old == nil) {
		// crash if old == nil.
		// could give a better message
		// but this is consistent with all the in-line checks
		// that the compiler inserts for other uses.
		*old = 0;
	}

	// new dope to old array
	ret.len = nel;
	ret.cap = nel;
	ret.array = old;

	FLUSH(&amp;ret);

	if(debug) {
		prints(&#34;runtime·slicearrayp: old=&#34;);
		runtime·printpointer(old);
		prints(&#34;; ret=&#34;);
		runtime·printslice(ret);
		prints(&#34;\n&#34;);
	}
}

void
runtime·printslice(Slice a)
{
	prints(&#34;[&#34;);
	runtime·printint(a.len);
	prints(&#34;/&#34;);
	runtime·printint(a.cap);
	prints(&#34;]&#34;);
	runtime·printpointer(a.array);
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/pswt.c</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/pswt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/6c/swt.c
// http://code.google.com/p/inferno-os/source/browse/utils/6c/swt.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the &#34;Software&#34;), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#include &#34;gc.h&#34;

int
swcmp(const void *a1, const void *a2)
{
	C1 *p1, *p2;

	p1 = (C1*)a1;
	p2 = (C1*)a2;
	if(p1-&gt;val &lt; p2-&gt;val)
		return -1;
	return p1-&gt;val &gt; p2-&gt;val;
}

void
doswit(Node *n)
{
	Case *c;
	C1 *q, *iq;
	int32 def, nc, i, isv;

	def = 0;
	nc = 0;
	isv = 0;
	for(c = cases; c-&gt;link != C; c = c-&gt;link) {
		if(c-&gt;def) {
			if(def)
				diag(n, &#34;more than one default in switch&#34;);
			def = c-&gt;label;
			continue;
		}
		isv |= c-&gt;isv;
		nc++;
	}
	if(isv &amp;&amp; !typev[n-&gt;type-&gt;etype])
		warn(n, &#34;32-bit switch expression with 64-bit case constant&#34;);

	iq = alloc(nc*sizeof(C1));
	q = iq;
	for(c = cases; c-&gt;link != C; c = c-&gt;link) {
		if(c-&gt;def)
			continue;
		q-&gt;label = c-&gt;label;
		if(isv)
			q-&gt;val = c-&gt;val;
		else
			q-&gt;val = (int32)c-&gt;val;	/* cast ensures correct value for 32-bit switch on 64-bit architecture */
		q++;
	}
	qsort(iq, nc, sizeof(C1), swcmp);
	if(debug[&#39;W&#39;])
	for(i=0; i&lt;nc; i++)
		print(&#34;case %2ld: = %.8llux\n&#34;, i, (vlong)iq[i].val);
	for(i=0; i&lt;nc-1; i++)
		if(iq[i].val == iq[i+1].val)
			diag(n, &#34;duplicate cases in switch %lld&#34;, (vlong)iq[i].val);
	if(def == 0) {
		def = breakpc;
		nbreak++;
	}
	swit1(iq, nc, def, n);
}

void
cas(void)
{
	Case *c;

	c = alloc(sizeof(*c));
	c-&gt;link = cases;
	cases = c;
}

int32
outlstring(ushort *s, int32 n)
{
	char buf[2];
	int c;
	int32 r;

	if(suppress)
		return nstring;
	while(nstring &amp; 1)
		outstring(&#34;&#34;, 1);
	r = nstring;
	while(n &gt; 0) {
		c = *s++;
		if(align(0, types[TCHAR], Aarg1)) {
			buf[0] = c&gt;&gt;8;
			buf[1] = c;
		} else {
			buf[0] = c;
			buf[1] = c&gt;&gt;8;
		}
		outstring(buf, 2);
		n -= sizeof(ushort);
	}
	return r;
}

void
nullwarn(Node *l, Node *r)
{
	warn(Z, &#34;result of operation not used&#34;);
	if(l != Z)
		cgen(l, Z);
	if(r != Z)
		cgen(r, Z);
}

void
ieeedtod(Ieee *ieee, double native)
{
	double fr, ho, f;
	int exp;

	if(native &lt; 0) {
		ieeedtod(ieee, -native);
		ieee-&gt;h |= 0x80000000L;
		return;
	}
	if(native == 0) {
		ieee-&gt;l = 0;
		ieee-&gt;h = 0;
		return;
	}
	fr = frexp(native, &amp;exp);
	f = 2097152L;		/* shouldnt use fp constants here */
	fr = modf(fr*f, &amp;ho);
	ieee-&gt;h = ho;
	ieee-&gt;h &amp;= 0xfffffL;
	ieee-&gt;h |= (exp+1022L) &lt;&lt; 20;
	f = 65536L;
	fr = modf(fr*f, &amp;ho);
	ieee-&gt;l = ho;
	ieee-&gt;l &lt;&lt;= 16;
	ieee-&gt;l |= (int32)(fr*f);
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

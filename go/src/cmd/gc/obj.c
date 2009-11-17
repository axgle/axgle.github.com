<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/obj.c</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/obj.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;go.h&#34;

/*
 * architecture-independent object file output
 */

void
dumpobj(void)
{
	bout = Bopen(outfile, OWRITE);
	if(bout == nil)
		fatal(&#34;cant open %s&#34;, outfile);

	Bprint(bout, &#34;%s\n&#34;, thestring);
	Bprint(bout, &#34;  exports automatically generated from\n&#34;);
	Bprint(bout, &#34;  %s in package \&#34;%s\&#34;\n&#34;, curio.infile, package);
	dumpexport();
	Bprint(bout, &#34;\n!\n&#34;);

	outhist(bout);

	// add nil plist w AEND to catch
	// auto-generated trampolines, data
	newplist();

	dumpglobls();
	dumptypestructs();
	dumpdata();
	dumpfuncs();

	Bterm(bout);
}

void
dumpglobls(void)
{
	Node *n;
	NodeList *l;

	// add globals
	for(l=externdcl; l; l=l-&gt;next) {
		n = l-&gt;n;
		if(n-&gt;op != ONAME)
			continue;

		if(n-&gt;type == T)
			fatal(&#34;external %#N nil type\n&#34;, n);
		if(n-&gt;class == PFUNC)
			continue;
		if(n-&gt;sym-&gt;package != package)
			continue;
		dowidth(n-&gt;type);

		// TODO(rsc): why is this not s/n-&gt;sym-&gt;def/n/ ?
		ggloblnod(n-&gt;sym-&gt;def, n-&gt;type-&gt;width);
	}
}

void
Bputdot(Biobuf *b)
{
	// put out middle dot ·
	Bputc(b, 0xc2);
	Bputc(b, 0xb7);
}

void
outhist(Biobuf *b)
{
	Hist *h;
	char *p, *q, *op;
	int n;

	for(h = hist; h != H; h = h-&gt;link) {
		p = h-&gt;name;
		op = 0;

		if(p &amp;&amp; p[0] != &#39;/&#39; &amp;&amp; h-&gt;offset == 0 &amp;&amp; pathname &amp;&amp; pathname[0] == &#39;/&#39;) {
			op = p;
			p = pathname;
		}

		while(p) {
			q = utfrune(p, &#39;/&#39;);
			if(q) {
				n = q-p;
				if(n == 0)
					n = 1;		// leading &#34;/&#34;
				q++;
			} else {
				n = strlen(p);
				q = 0;
			}
			if(n)
				zfile(b, p, n);
			p = q;
			if(p == 0 &amp;&amp; op) {
				p = op;
				op = 0;
			}
		}

		zhist(b, h-&gt;line, h-&gt;offset);
	}
}

void
ieeedtod(uint64 *ieee, double native)
{
	double fr, ho, f;
	int exp;
	uint32 h, l;

	if(native &lt; 0) {
		ieeedtod(ieee, -native);
		*ieee |= 1ULL&lt;&lt;63;
		return;
	}
	if(native == 0) {
		*ieee = 0;
		return;
	}
	fr = frexp(native, &amp;exp);
	f = 2097152L;		/* shouldnt use fp constants here */
	fr = modf(fr*f, &amp;ho);
	h = ho;
	h &amp;= 0xfffffL;
	h |= (exp+1022L) &lt;&lt; 20;
	f = 65536L;
	fr = modf(fr*f, &amp;ho);
	l = ho;
	l &lt;&lt;= 16;
	l |= (int32)(fr*f);
	*ieee = ((uint64)h &lt;&lt; 32) | l;
}

int
duint8(Sym *s, int off, uint8 v)
{
	return duintxx(s, off, v, 1);
}

int
duint16(Sym *s, int off, uint16 v)
{
	return duintxx(s, off, v, 2);
}

int
duint32(Sym *s, int off, uint32 v)
{
	return duintxx(s, off, v, 4);
}

int
duint64(Sym *s, int off, uint64 v)
{
	return duintxx(s, off, v, 8);
}

int
duintptr(Sym *s, int off, uint64 v)
{
	return duintxx(s, off, v, widthptr);
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/print.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/print.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;go.h&#34;

enum
{
	PFIXME = 0,
	PCHAN = 0,
};

void
exprlistfmt(Fmt *f, NodeList *l)
{
	for(; l; l=l-&gt;next) {
		exprfmt(f, l-&gt;n, 0);
		if(l-&gt;next)
			fmtprint(f, &#34;, &#34;);
	}
}

void
exprfmt(Fmt *f, Node *n, int prec)
{
	int nprec;

	nprec = 0;
	if(n == nil) {
		fmtprint(f, &#34;&lt;nil&gt;&#34;);
		return;
	}

	switch(n-&gt;op) {
	case ONAME:
	case ONONAME:
	case OPACK:
	case OLITERAL:
	case ODOT:
	case ODOTPTR:
	case ODOTINTER:
	case ODOTMETH:
	case OARRAYBYTESTR:
	case OCAP:
	case OCLOSE:
	case OCLOSED:
	case OLEN:
	case OMAKE:
	case ONEW:
	case OPANIC:
	case OPANICN:
	case OPRINT:
	case OPRINTN:
	case OCALL:
	case OCONV:
	case OCONVNOP:
	case OCONVSLICE:
	case OCONVIFACE:
	case OMAKESLICE:
	case ORUNESTR:
	case OADDR:
	case OCOM:
	case OIND:
	case OMINUS:
	case ONOT:
	case OPLUS:
	case ORECV:
		nprec = 7;
		break;

	case OMUL:
	case ODIV:
	case OMOD:
	case OLSH:
	case ORSH:
	case OAND:
	case OANDNOT:
		nprec = 6;
		break;

	case OADD:
	case OSUB:
	case OOR:
	case OXOR:
		nprec = 5;
		break;

	case OEQ:
	case OLT:
	case OLE:
	case OGE:
	case OGT:
	case ONE:
		nprec = 4;
		break;

	case OSEND:
		nprec = 3;
		break;

	case OANDAND:
		nprec = 2;
		break;

	case OOROR:
		nprec = 1;
		break;
	}

	if(prec &gt; nprec)
		fmtprint(f, &#34;(&#34;);

	switch(n-&gt;op) {
	default:
	bad:
		fmtprint(f, &#34;(node %O)&#34;, n-&gt;op);
		break;

	case OLITERAL:
		switch(n-&gt;val.ctype) {
		default:
			goto bad;
		case CTINT:
			fmtprint(f, &#34;%B&#34;, n-&gt;val.u.xval);
			break;
		case CTBOOL:
			if(n-&gt;val.u.bval)
				fmtprint(f, &#34;true&#34;);
			else
				fmtprint(f, &#34;false&#34;);
			break;
		case CTFLT:
			fmtprint(f, &#34;%.17g&#34;, mpgetflt(n-&gt;val.u.fval));
			break;
		case CTSTR:
			fmtprint(f, &#34;\&#34;%Z\&#34;&#34;, n-&gt;val.u.sval);
			break;
		case CTNIL:
			fmtprint(f, &#34;nil&#34;);
			break;
		}
		break;

	case ONAME:
	case OPACK:
	case ONONAME:
		fmtprint(f, &#34;%S&#34;, n-&gt;sym);
		break;

	case OTYPE:
		fmtprint(f, &#34;%T&#34;, n-&gt;type);
		break;

	case OTARRAY:
		fmtprint(f, &#34;[]&#34;);
		exprfmt(f, n-&gt;left, PFIXME);
		break;

	case OTMAP:
		fmtprint(f, &#34;map[&#34;);
		exprfmt(f, n-&gt;left, 0);
		fmtprint(f, &#34;] &#34;);
		exprfmt(f, n-&gt;right, 0);
		break;

	case OTCHAN:
		if(n-&gt;etype == Crecv)
			fmtprint(f, &#34;&lt;-&#34;);
		fmtprint(f, &#34;chan&#34;);
		if(n-&gt;etype == Csend) {
			fmtprint(f, &#34;&lt;- &#34;);
			exprfmt(f, n-&gt;left, 0);
		} else {
			fmtprint(f, &#34; &#34;);
			exprfmt(f, n-&gt;left, PCHAN);
		}
		break;

	case OTSTRUCT:
		fmtprint(f, &#34;&lt;struct&gt;&#34;);
		break;

	case OTINTER:
		fmtprint(f, &#34;&lt;inter&gt;&#34;);
		break;

	case OTFUNC:
		fmtprint(f, &#34;&lt;func&gt;&#34;);
		break;

	case OAS:
		exprfmt(f, n-&gt;left, 0);
		fmtprint(f, &#34; = &#34;);
		exprfmt(f, n-&gt;right, 0);
		break;

	case OASOP:
		exprfmt(f, n-&gt;left, 0);
		fmtprint(f, &#34; %#O= &#34;, n-&gt;etype);
		exprfmt(f, n-&gt;right, 0);
		break;

	case OADD:
	case OANDAND:
	case OANDNOT:
	case ODIV:
	case OEQ:
	case OGE:
	case OGT:
	case OLE:
	case OLT:
	case OLSH:
	case OMOD:
	case OMUL:
	case ONE:
	case OOR:
	case OOROR:
	case ORSH:
	case OSEND:
	case OSUB:
	case OXOR:
		exprfmt(f, n-&gt;left, nprec);
		fmtprint(f, &#34; %#O &#34;, n-&gt;op);
		exprfmt(f, n-&gt;right, nprec+1);
		break;

	case OADDR:
	case OCOM:
	case OIND:
	case OMINUS:
	case ONOT:
	case OPLUS:
	case ORECV:
		fmtprint(f, &#34;%#O&#34;, n-&gt;op);
		if((n-&gt;op == OMINUS || n-&gt;op == OPLUS) &amp;&amp; n-&gt;left-&gt;op == n-&gt;op)
			fmtprint(f, &#34; &#34;);
		exprfmt(f, n-&gt;left, 0);
		break;

	case OCOMPLIT:
		fmtprint(f, &#34;&lt;compos&gt;&#34;);
		break;

	case ODOT:
	case ODOTPTR:
	case ODOTINTER:
	case ODOTMETH:
		exprfmt(f, n-&gt;left, 7);
		if(n-&gt;right == N || n-&gt;right-&gt;sym == S)
			fmtprint(f, &#34;.&lt;nil&gt;&#34;);
		else
			fmtprint(f, &#34;.%s&#34;, n-&gt;right-&gt;sym-&gt;name);
		break;

	case ODOTTYPE:
		exprfmt(f, n-&gt;left, 7);
		fmtprint(f, &#34;.(&#34;);
		exprfmt(f, n-&gt;right, 0);
		fmtprint(f, &#34;)&#34;);
		break;

	case OINDEX:
	case OINDEXMAP:
	case OINDEXSTR:
		exprfmt(f, n-&gt;left, 7);
		fmtprint(f, &#34;[&#34;);
		exprfmt(f, n-&gt;right, 0);
		fmtprint(f, &#34;]&#34;);
		break;

	case OSLICE:
		exprfmt(f, n-&gt;left, 7);
		fmtprint(f, &#34;[&#34;);
		exprfmt(f, n-&gt;right-&gt;left, 0);
		fmtprint(f, &#34;:&#34;);
		exprfmt(f, n-&gt;right-&gt;right, 0);
		fmtprint(f, &#34;]&#34;);
		break;

	case OCALL:
	case OCALLFUNC:
	case OCALLINTER:
	case OCALLMETH:
		exprfmt(f, n-&gt;left, 7);
		fmtprint(f, &#34;(&#34;);
		exprlistfmt(f, n-&gt;list);
		fmtprint(f, &#34;)&#34;);
		break;

	case OCONV:
	case OCONVNOP:
	case OCONVSLICE:
	case OCONVIFACE:
	case OARRAYBYTESTR:
	case ORUNESTR:
		if(n-&gt;type == T || n-&gt;type-&gt;sym == S)
			fmtprint(f, &#34;(%T)(&#34;, n-&gt;type);
		else
			fmtprint(f, &#34;%T(&#34;, n-&gt;type);
		exprfmt(f, n-&gt;left, 0);
		fmtprint(f, &#34;)&#34;);
		break;

	case OCAP:
	case OCLOSE:
	case OCLOSED:
	case OLEN:
	case OMAKE:
	case ONEW:
	case OPANIC:
	case OPANICN:
	case OPRINT:
	case OPRINTN:
		fmtprint(f, &#34;%#O(&#34;, n-&gt;op);
		if(n-&gt;left)
			exprfmt(f, n-&gt;left, 0);
		else
			exprlistfmt(f, n-&gt;list);
		fmtprint(f, &#34;)&#34;);
		break;

	case OMAKESLICE:
		fmtprint(f, &#34;make(%#T, &#34;, n-&gt;type);
		exprfmt(f, n-&gt;left, 0);
		if(count(n-&gt;list) &gt; 2) {
			fmtprint(f, &#34;, &#34;);
			exprfmt(f, n-&gt;right, 0);
		}
		fmtprint(f, &#34;)&#34;);
		break;

	case OMAKEMAP:
		fmtprint(f, &#34;make(%#T)&#34;, n-&gt;type);
		break;

	case OMAPLIT:
		fmtprint(f, &#34;map literal&#34;);
	}

	if(prec &gt; nprec)
		fmtprint(f, &#34;)&#34;);
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

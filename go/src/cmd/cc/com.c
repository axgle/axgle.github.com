<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/com.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/cc/com.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/com.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/com.c
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

#include &#34;cc.h&#34;

int compar(Node*, int);

void
complex(Node *n)
{

	if(n == Z)
		return;

	nearln = n-&gt;lineno;
	if(debug[&#39;t&#39;])
		if(n-&gt;op != OCONST)
			prtree(n, &#34;pre complex&#34;);
	if(tcom(n))
		return;
	if(debug[&#39;t&#39;])
		if(n-&gt;op != OCONST)
			prtree(n, &#34;t complex&#34;);
	ccom(n);
	if(debug[&#39;t&#39;])
		if(n-&gt;op != OCONST)
			prtree(n, &#34;c complex&#34;);
	acom(n);
	if(debug[&#39;t&#39;])
		if(n-&gt;op != OCONST)
			prtree(n, &#34;a complex&#34;);
	xcom(n);
	if(debug[&#39;t&#39;])
		if(n-&gt;op != OCONST)
			prtree(n, &#34;x complex&#34;);
}

/*
 * evaluate types
 * evaluate lvalues (addable == 1)
 */
enum
{
	ADDROF	= 1&lt;&lt;0,
	CASTOF	= 1&lt;&lt;1,
	ADDROP	= 1&lt;&lt;2,
};

int
tcom(Node *n)
{

	return tcomo(n, ADDROF);
}

int
tcomo(Node *n, int f)
{
	Node *l, *r;
	Type *t;
	int o;

	if(n == Z) {
		diag(Z, &#34;Z in tcom&#34;);
		errorexit();
	}
	n-&gt;addable = 0;
	l = n-&gt;left;
	r = n-&gt;right;

	switch(n-&gt;op) {
	default:
		diag(n, &#34;unknown op in type complex: %O&#34;, n-&gt;op);
		goto bad;

	case ODOTDOT:
		/*
		 * tcom has already been called on this subtree
		 */
		*n = *n-&gt;left;
		if(n-&gt;type == T)
			goto bad;
		break;

	case OCAST:
		if(n-&gt;type == T)
			break;
		if(n-&gt;type-&gt;width == types[TLONG]-&gt;width) {
			if(tcomo(l, ADDROF|CASTOF))
				goto bad;
		} else
			if(tcom(l))
				goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, n-&gt;type, tcast))
			goto bad;
		break;

	case ORETURN:
		if(l == Z) {
			if(n-&gt;type-&gt;etype != TVOID)
				warn(n, &#34;null return of a typed function&#34;);
			break;
		}
		if(tcom(l))
			goto bad;
		typeext(n-&gt;type, l);
		if(tcompat(n, n-&gt;type, l-&gt;type, tasign))
			break;
		constas(n, n-&gt;type, l-&gt;type);
		if(!sametype(n-&gt;type, l-&gt;type)) {
			l = new1(OCAST, l, Z);
			l-&gt;type = n-&gt;type;
			n-&gt;left = l;
		}
		break;

	case OASI:	/* same as as, but no test for const */
		n-&gt;op = OAS;
		o = tcom(l);
		if(o | tcom(r))
			goto bad;

		typeext(l-&gt;type, r);
		if(tlvalue(l) || tcompat(n, l-&gt;type, r-&gt;type, tasign))
			goto bad;
		if(!sametype(l-&gt;type, r-&gt;type)) {
			r = new1(OCAST, r, Z);
			r-&gt;type = l-&gt;type;
			n-&gt;right = r;
		}
		n-&gt;type = l-&gt;type;
		break;

	case OAS:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(tlvalue(l))
			goto bad;
		if(isfunct(n))
			break;
		typeext(l-&gt;type, r);
		if(tcompat(n, l-&gt;type, r-&gt;type, tasign))
			goto bad;
		constas(n, l-&gt;type, r-&gt;type);
		if(!sametype(l-&gt;type, r-&gt;type)) {
			r = new1(OCAST, r, Z);
			r-&gt;type = l-&gt;type;
			n-&gt;right = r;
		}
		n-&gt;type = l-&gt;type;
		break;

	case OASADD:
	case OASSUB:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(tlvalue(l))
			goto bad;
		if(isfunct(n))
			break;
		typeext1(l-&gt;type, r);
		if(tcompat(n, l-&gt;type, r-&gt;type, tasadd))
			goto bad;
		constas(n, l-&gt;type, r-&gt;type);
		t = l-&gt;type;
		arith(n, 0);
		while(n-&gt;left-&gt;op == OCAST)
			n-&gt;left = n-&gt;left-&gt;left;
		if(!sametype(t, n-&gt;type) &amp;&amp; !mixedasop(t, n-&gt;type)) {
			r = new1(OCAST, n-&gt;right, Z);
			r-&gt;type = t;
			n-&gt;right = r;
			n-&gt;type = t;
		}
		break;

	case OASMUL:
	case OASLMUL:
	case OASDIV:
	case OASLDIV:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(tlvalue(l))
			goto bad;
		if(isfunct(n))
			break;
		typeext1(l-&gt;type, r);
		if(tcompat(n, l-&gt;type, r-&gt;type, tmul))
			goto bad;
		constas(n, l-&gt;type, r-&gt;type);
		t = l-&gt;type;
		arith(n, 0);
		while(n-&gt;left-&gt;op == OCAST)
			n-&gt;left = n-&gt;left-&gt;left;
		if(!sametype(t, n-&gt;type) &amp;&amp; !mixedasop(t, n-&gt;type)) {
			r = new1(OCAST, n-&gt;right, Z);
			r-&gt;type = t;
			n-&gt;right = r;
			n-&gt;type = t;
		}
		if(typeu[n-&gt;type-&gt;etype]) {
			if(n-&gt;op == OASDIV)
				n-&gt;op = OASLDIV;
			if(n-&gt;op == OASMUL)
				n-&gt;op = OASLMUL;
		}
		break;

	case OASLSHR:
	case OASASHR:
	case OASASHL:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(tlvalue(l))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, r-&gt;type, tand))
			goto bad;
		n-&gt;type = l-&gt;type;
		if(typeu[n-&gt;type-&gt;etype]) {
			if(n-&gt;op == OASASHR)
				n-&gt;op = OASLSHR;
		}
		break;

	case OASMOD:
	case OASLMOD:
	case OASOR:
	case OASAND:
	case OASXOR:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(tlvalue(l))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, r-&gt;type, tand))
			goto bad;
		t = l-&gt;type;
		arith(n, 0);
		while(n-&gt;left-&gt;op == OCAST)
			n-&gt;left = n-&gt;left-&gt;left;
		if(!sametype(t, n-&gt;type) &amp;&amp; !mixedasop(t, n-&gt;type)) {
			r = new1(OCAST, n-&gt;right, Z);
			r-&gt;type = t;
			n-&gt;right = r;
			n-&gt;type = t;
		}
		if(typeu[n-&gt;type-&gt;etype]) {
			if(n-&gt;op == OASMOD)
				n-&gt;op = OASLMOD;
		}
		break;

	case OPREINC:
	case OPREDEC:
	case OPOSTINC:
	case OPOSTDEC:
		if(tcom(l))
			goto bad;
		if(tlvalue(l))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, types[TINT], tadd))
			goto bad;
		n-&gt;type = l-&gt;type;
		if(n-&gt;type-&gt;etype == TIND)
		if(n-&gt;type-&gt;link-&gt;width &lt; 1)
			diag(n, &#34;inc/dec of a void pointer&#34;);
		break;

	case OEQ:
	case ONE:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(isfunct(n))
			break;
		typeext(l-&gt;type, r);
		typeext(r-&gt;type, l);
		if(tcompat(n, l-&gt;type, r-&gt;type, trel))
			goto bad;
		arith(n, 0);
		n-&gt;type = types[TINT];
		break;

	case OLT:
	case OGE:
	case OGT:
	case OLE:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(isfunct(n))
			break;
		typeext1(l-&gt;type, r);
		typeext1(r-&gt;type, l);
		if(tcompat(n, l-&gt;type, r-&gt;type, trel))
			goto bad;
		arith(n, 0);
		if(typeu[n-&gt;type-&gt;etype])
			n-&gt;op = logrel[relindex(n-&gt;op)];
		n-&gt;type = types[TINT];
		break;

	case OCOND:
		o = tcom(l);
		o |= tcom(r-&gt;left);
		if(o | tcom(r-&gt;right))
			goto bad;
		if(r-&gt;right-&gt;type-&gt;etype == TIND &amp;&amp; vconst(r-&gt;left) == 0) {
			r-&gt;left-&gt;type = r-&gt;right-&gt;type;
			r-&gt;left-&gt;vconst = 0;
		}
		if(r-&gt;left-&gt;type-&gt;etype == TIND &amp;&amp; vconst(r-&gt;right) == 0) {
			r-&gt;right-&gt;type = r-&gt;left-&gt;type;
			r-&gt;right-&gt;vconst = 0;
		}
		if(sametype(r-&gt;right-&gt;type, r-&gt;left-&gt;type)) {
			r-&gt;type = r-&gt;right-&gt;type;
			n-&gt;type = r-&gt;type;
			break;
		}
		if(tcompat(r, r-&gt;left-&gt;type, r-&gt;right-&gt;type, trel))
			goto bad;
		arith(r, 0);
		n-&gt;type = r-&gt;type;
		break;

	case OADD:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, r-&gt;type, tadd))
			goto bad;
		arith(n, 1);
		break;

	case OSUB:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, r-&gt;type, tsub))
			goto bad;
		arith(n, 1);
		break;

	case OMUL:
	case OLMUL:
	case ODIV:
	case OLDIV:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, r-&gt;type, tmul))
			goto bad;
		arith(n, 1);
		if(typeu[n-&gt;type-&gt;etype]) {
			if(n-&gt;op == ODIV)
				n-&gt;op = OLDIV;
			if(n-&gt;op == OMUL)
				n-&gt;op = OLMUL;
		}
		break;

	case OLSHR:
	case OASHL:
	case OASHR:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, r-&gt;type, tand))
			goto bad;
		n-&gt;right = Z;
		arith(n, 1);
		n-&gt;right = new1(OCAST, r, Z);
		n-&gt;right-&gt;type = types[TINT];
		if(typeu[n-&gt;type-&gt;etype])
			if(n-&gt;op == OASHR)
				n-&gt;op = OLSHR;
		break;

	case OAND:
	case OOR:
	case OXOR:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, r-&gt;type, tand))
			goto bad;
		arith(n, 1);
		break;

	case OMOD:
	case OLMOD:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, l-&gt;type, r-&gt;type, tand))
			goto bad;
		arith(n, 1);
		if(typeu[n-&gt;type-&gt;etype])
			n-&gt;op = OLMOD;
		break;

	case OPOS:
		if(tcom(l))
			goto bad;
		if(isfunct(n))
			break;

		r = l;
		l = new(OCONST, Z, Z);
		l-&gt;vconst = 0;
		l-&gt;type = types[TINT];
		n-&gt;op = OADD;
		n-&gt;right = r;
		n-&gt;left = l;

		if(tcom(l))
			goto bad;
		if(tcompat(n, l-&gt;type, r-&gt;type, tsub))
			goto bad;
		arith(n, 1);
		break;

	case ONEG:
		if(tcom(l))
			goto bad;
		if(isfunct(n))
			break;

		if(!machcap(n)) {
			r = l;
			l = new(OCONST, Z, Z);
			l-&gt;vconst = 0;
			l-&gt;type = types[TINT];
			n-&gt;op = OSUB;
			n-&gt;right = r;
			n-&gt;left = l;

			if(tcom(l))
				goto bad;
			if(tcompat(n, l-&gt;type, r-&gt;type, tsub))
				goto bad;
		}
		arith(n, 1);
		break;

	case OCOM:
		if(tcom(l))
			goto bad;
		if(isfunct(n))
			break;

		if(!machcap(n)) {
			r = l;
			l = new(OCONST, Z, Z);
			l-&gt;vconst = -1;
			l-&gt;type = types[TINT];
			n-&gt;op = OXOR;
			n-&gt;right = r;
			n-&gt;left = l;

			if(tcom(l))
				goto bad;
			if(tcompat(n, l-&gt;type, r-&gt;type, tand))
				goto bad;
		}
		arith(n, 1);
		break;

	case ONOT:
		if(tcom(l))
			goto bad;
		if(isfunct(n))
			break;
		if(tcompat(n, T, l-&gt;type, tnot))
			goto bad;
		n-&gt;type = types[TINT];
		break;

	case OANDAND:
	case OOROR:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		if(tcompat(n, T, l-&gt;type, tnot) |
		   tcompat(n, T, r-&gt;type, tnot))
			goto bad;
		n-&gt;type = types[TINT];
		break;

	case OCOMMA:
		o = tcom(l);
		if(o | tcom(r))
			goto bad;
		n-&gt;type = r-&gt;type;
		break;


	case OSIGN:	/* extension signof(type) returns a hash */
		if(l != Z) {
			if(l-&gt;op != OSTRING &amp;&amp; l-&gt;op != OLSTRING)
				if(tcomo(l, 0))
					goto bad;
			if(l-&gt;op == OBIT) {
				diag(n, &#34;signof bitfield&#34;);
				goto bad;
			}
			n-&gt;type = l-&gt;type;
		}
		if(n-&gt;type == T)
			goto bad;
		if(n-&gt;type-&gt;width &lt; 0) {
			diag(n, &#34;signof undefined type&#34;);
			goto bad;
		}
		n-&gt;op = OCONST;
		n-&gt;left = Z;
		n-&gt;right = Z;
		n-&gt;vconst = convvtox(signature(n-&gt;type), TULONG);
		n-&gt;type = types[TULONG];
		break;

	case OSIZE:
		if(l != Z) {
			if(l-&gt;op != OSTRING &amp;&amp; l-&gt;op != OLSTRING)
				if(tcomo(l, 0))
					goto bad;
			if(l-&gt;op == OBIT) {
				diag(n, &#34;sizeof bitfield&#34;);
				goto bad;
			}
			n-&gt;type = l-&gt;type;
		}
		if(n-&gt;type == T)
			goto bad;
		if(n-&gt;type-&gt;width &lt;= 0) {
			diag(n, &#34;sizeof undefined type&#34;);
			goto bad;
		}
		if(n-&gt;type-&gt;etype == TFUNC) {
			diag(n, &#34;sizeof function&#34;);
			goto bad;
		}
		n-&gt;op = OCONST;
		n-&gt;left = Z;
		n-&gt;right = Z;
		n-&gt;vconst = convvtox(n-&gt;type-&gt;width, TINT);
		n-&gt;type = types[TINT];
		break;

	case OFUNC:
		o = tcomo(l, 0);
		if(o)
			goto bad;
		if(l-&gt;type-&gt;etype == TIND &amp;&amp; l-&gt;type-&gt;link-&gt;etype == TFUNC) {
			l = new1(OIND, l, Z);
			l-&gt;type = l-&gt;left-&gt;type-&gt;link;
			n-&gt;left = l;
		}
		if(tcompat(n, T, l-&gt;type, tfunct))
			goto bad;
		if(o | tcoma(l, r, l-&gt;type-&gt;down, 1))
			goto bad;
		n-&gt;type = l-&gt;type-&gt;link;
		if(!debug[&#39;B&#39;])
			if(l-&gt;type-&gt;down == T || l-&gt;type-&gt;down-&gt;etype == TOLD) {
				nerrors--;
				diag(n, &#34;function args not checked: %F&#34;, l);
			}
		dpcheck(n);
		break;

	case ONAME:
		if(n-&gt;type == T) {
			diag(n, &#34;name not declared: %F&#34;, n);
			goto bad;
		}
		if(n-&gt;type-&gt;etype == TENUM) {
			n-&gt;op = OCONST;
			n-&gt;type = n-&gt;sym-&gt;tenum;
			if(!typefd[n-&gt;type-&gt;etype])
				n-&gt;vconst = n-&gt;sym-&gt;vconst;
			else
				n-&gt;fconst = n-&gt;sym-&gt;fconst;
			break;
		}
		n-&gt;addable = 1;
		if(n-&gt;class == CEXREG) {
			n-&gt;op = OREGISTER;
			// on 386, &#34;extern register&#34; generates
			// memory references relative to the
			// fs segment.
			if(thechar == &#39;8&#39;)	// [sic]
				n-&gt;op = OEXREG;
			n-&gt;reg = n-&gt;sym-&gt;offset;
			n-&gt;xoffset = 0;
			break;
		}
		break;

	case OLSTRING:
		if(n-&gt;type-&gt;link != types[TUSHORT]) {
			o = outstring(0, 0);
			while(o &amp; 3) {
				ushort a[1];
				a[0] = 0;
				outlstring(a, sizeof(ushort));
				o = outlstring(0, 0);
			}
		}
		n-&gt;op = ONAME;
		n-&gt;xoffset = outlstring(n-&gt;rstring, n-&gt;type-&gt;width);
		n-&gt;addable = 1;
		break;

	case OSTRING:
		if(n-&gt;type-&gt;link != types[TCHAR]) {
			o = outstring(0, 0);
			while(o &amp; 3) {
				outstring(&#34;&#34;, 1);
				o = outstring(0, 0);
			}
		}
		n-&gt;op = ONAME;
		n-&gt;xoffset = outstring(n-&gt;cstring, n-&gt;type-&gt;width);
		n-&gt;addable = 1;
		break;

	case OCONST:
		break;

	case ODOT:
		if(tcom(l))
			goto bad;
		if(tcompat(n, T, l-&gt;type, tdot))
			goto bad;
		if(tcomd(n))
			goto bad;
		break;

	case OADDR:
		if(tcomo(l, ADDROP))
			goto bad;
		if(tlvalue(l))
			goto bad;
		if(l-&gt;type-&gt;nbits) {
			diag(n, &#34;address of a bit field&#34;);
			goto bad;
		}
		if(l-&gt;op == OREGISTER) {
			diag(n, &#34;address of a register&#34;);
			goto bad;
		}
		n-&gt;type = typ(TIND, l-&gt;type);
		n-&gt;type-&gt;width = types[TIND]-&gt;width;
		break;

	case OIND:
		if(tcom(l))
			goto bad;
		if(tcompat(n, T, l-&gt;type, tindir))
			goto bad;
		n-&gt;type = l-&gt;type-&gt;link;
		n-&gt;addable = 1;
		break;

	case OSTRUCT:
		if(tcomx(n))
			goto bad;
		break;
	}
	t = n-&gt;type;
	if(t == T)
		goto bad;
	if(t-&gt;width &lt; 0) {
		snap(t);
		if(t-&gt;width &lt; 0) {
			if(typesu[t-&gt;etype] &amp;&amp; t-&gt;tag)
				diag(n, &#34;structure not fully declared %s&#34;, t-&gt;tag-&gt;name);
			else
				diag(n, &#34;structure not fully declared&#34;);
			goto bad;
		}
	}
	if(typeaf[t-&gt;etype]) {
		if(f &amp; ADDROF)
			goto addaddr;
		if(f &amp; ADDROP)
			warn(n, &#34;address of array/func ignored&#34;);
	}
	return 0;

addaddr:
	if(tlvalue(n))
		goto bad;
	l = new1(OXXX, Z, Z);
	*l = *n;
	n-&gt;op = OADDR;
	if(l-&gt;type-&gt;etype == TARRAY)
		l-&gt;type = l-&gt;type-&gt;link;
	n-&gt;left = l;
	n-&gt;right = Z;
	n-&gt;addable = 0;
	n-&gt;type = typ(TIND, l-&gt;type);
	n-&gt;type-&gt;width = types[TIND]-&gt;width;
	return 0;

bad:
	n-&gt;type = T;
	return 1;
}

int
tcoma(Node *l, Node *n, Type *t, int f)
{
	Node *n1;
	int o;

	if(t != T)
	if(t-&gt;etype == TOLD || t-&gt;etype == TDOT)	/* .../old in prototype */
		t = T;
	if(n == Z) {
		if(t != T &amp;&amp; !sametype(t, types[TVOID])) {
			diag(n, &#34;not enough function arguments: %F&#34;, l);
			return 1;
		}
		return 0;
	}
	if(n-&gt;op == OLIST) {
		o = tcoma(l, n-&gt;left, t, 0);
		if(t != T) {
			t = t-&gt;down;
			if(t == T)
				t = types[TVOID];
		}
		return o | tcoma(l, n-&gt;right, t, 1);
	}
	if(f &amp;&amp; t != T)
		tcoma(l, Z, t-&gt;down, 0);
	if(tcom(n) || tcompat(n, T, n-&gt;type, targ))
		return 1;
	if(sametype(t, types[TVOID])) {
		diag(n, &#34;too many function arguments: %F&#34;, l);
		return 1;
	}
	if(t != T) {
		typeext(t, n);
		if(stcompat(nodproto, t, n-&gt;type, tasign)) {
			diag(l, &#34;argument prototype mismatch \&#34;%T\&#34; for \&#34;%T\&#34;: %F&#34;,
				n-&gt;type, t, l);
			return 1;
		}
		switch(t-&gt;etype) {
		case TCHAR:
		case TSHORT:
			t = types[TINT];
			break;

		case TUCHAR:
		case TUSHORT:
			t = types[TUINT];
			break;
		}
	} else
	switch(n-&gt;type-&gt;etype)
	{
	case TCHAR:
	case TSHORT:
		t = types[TINT];
		break;

	case TUCHAR:
	case TUSHORT:
		t = types[TUINT];
		break;

	case TFLOAT:
		t = types[TDOUBLE];
	}
	if(t != T &amp;&amp; !sametype(t, n-&gt;type)) {
		n1 = new1(OXXX, Z, Z);
		*n1 = *n;
		n-&gt;op = OCAST;
		n-&gt;left = n1;
		n-&gt;right = Z;
		n-&gt;type = t;
		n-&gt;addable = 0;
	}
	return 0;
}

int
tcomd(Node *n)
{
	Type *t;
	int32 o;

	o = 0;
	t = dotsearch(n-&gt;sym, n-&gt;left-&gt;type-&gt;link, n, &amp;o);
	if(t == T) {
		diag(n, &#34;not a member of struct/union: %F&#34;, n);
		return 1;
	}
	makedot(n, t, o);
	return 0;
}

int
tcomx(Node *n)
{
	Type *t;
	Node *l, *r, **ar, **al;
	int e;

	e = 0;
	if(n-&gt;type-&gt;etype != TSTRUCT) {
		diag(n, &#34;constructor must be a structure&#34;);
		return 1;
	}
	l = invert(n-&gt;left);
	n-&gt;left = l;
	al = &amp;n-&gt;left;
	for(t = n-&gt;type-&gt;link; t != T; t = t-&gt;down) {
		if(l == Z) {
			diag(n, &#34;constructor list too short&#34;);
			return 1;
		}
		if(l-&gt;op == OLIST) {
			r = l-&gt;left;
			ar = &amp;l-&gt;left;
			al = &amp;l-&gt;right;
			l = l-&gt;right;
		} else {
			r = l;
			ar = al;
			l = Z;
		}
		if(tcom(r))
			e++;
		typeext(t, r);
		if(tcompat(n, t, r-&gt;type, tasign))
			e++;
		constas(n, t, r-&gt;type);
		if(!e &amp;&amp; !sametype(t, r-&gt;type)) {
			r = new1(OCAST, r, Z);
			r-&gt;type = t;
			*ar = r;
		}
	}
	if(l != Z) {
		diag(n, &#34;constructor list too long&#34;);
		return 1;
	}
	return e;
}

int
tlvalue(Node *n)
{

	if(!n-&gt;addable) {
		diag(n, &#34;not an l-value&#34;);
		return 1;
	}
	return 0;
}

/*
 *	general rewrite
 *	(IND(ADDR x)) ==&gt; x
 *	(ADDR(IND x)) ==&gt; x
 *	remove some zero operands
 *	remove no op casts
 *	evaluate constants
 */
void
ccom(Node *n)
{
	Node *l, *r;
	int t;

loop:
	if(n == Z)
		return;
	l = n-&gt;left;
	r = n-&gt;right;
	switch(n-&gt;op) {

	case OAS:
	case OASXOR:
	case OASAND:
	case OASOR:
	case OASMOD:
	case OASLMOD:
	case OASLSHR:
	case OASASHR:
	case OASASHL:
	case OASDIV:
	case OASLDIV:
	case OASMUL:
	case OASLMUL:
	case OASSUB:
	case OASADD:
		ccom(l);
		ccom(r);
		if(n-&gt;op == OASLSHR || n-&gt;op == OASASHR || n-&gt;op == OASASHL)
		if(r-&gt;op == OCONST) {
			t = n-&gt;type-&gt;width * 8;	/* bits per byte */
			if(r-&gt;vconst &gt;= t || r-&gt;vconst &lt; 0)
				warn(n, &#34;stupid shift: %lld&#34;, r-&gt;vconst);
		}
		break;

	case OCAST:
		ccom(l);
		if(l-&gt;op == OCONST) {
			evconst(n);
			if(n-&gt;op == OCONST)
				break;
		}
		if(nocast(l-&gt;type, n-&gt;type)) {
			l-&gt;type = n-&gt;type;
			*n = *l;
		}
		break;

	case OCOND:
		ccom(l);
		ccom(r);
		if(l-&gt;op == OCONST)
			if(vconst(l) == 0)
				*n = *r-&gt;right;
			else
				*n = *r-&gt;left;
		break;

	case OREGISTER:
	case OINDREG:
	case OCONST:
	case ONAME:
		break;

	case OADDR:
		ccom(l);
		l-&gt;etype = TVOID;
		if(l-&gt;op == OIND) {
			l-&gt;left-&gt;type = n-&gt;type;
			*n = *l-&gt;left;
			break;
		}
		goto common;

	case OIND:
		ccom(l);
		if(l-&gt;op == OADDR) {
			l-&gt;left-&gt;type = n-&gt;type;
			*n = *l-&gt;left;
			break;
		}
		goto common;

	case OEQ:
	case ONE:

	case OLE:
	case OGE:
	case OLT:
	case OGT:

	case OLS:
	case OHS:
	case OLO:
	case OHI:
		ccom(l);
		ccom(r);
		if(compar(n, 0) || compar(n, 1))
			break;
		relcon(l, r);
		relcon(r, l);
		goto common;

	case OASHR:
	case OASHL:
	case OLSHR:
		ccom(l);
		if(vconst(l) == 0 &amp;&amp; !side(r)) {
			*n = *l;
			break;
		}
		ccom(r);
		if(vconst(r) == 0) {
			*n = *l;
			break;
		}
		if(r-&gt;op == OCONST) {
			t = n-&gt;type-&gt;width * 8;	/* bits per byte */
			if(r-&gt;vconst &gt;= t || r-&gt;vconst &lt;= -t)
				warn(n, &#34;stupid shift: %lld&#34;, r-&gt;vconst);
		}
		goto common;

	case OMUL:
	case OLMUL:
		ccom(l);
		t = vconst(l);
		if(t == 0 &amp;&amp; !side(r)) {
			*n = *l;
			break;
		}
		if(t == 1) {
			*n = *r;
			goto loop;
		}
		ccom(r);
		t = vconst(r);
		if(t == 0 &amp;&amp; !side(l)) {
			*n = *r;
			break;
		}
		if(t == 1) {
			*n = *l;
			break;
		}
		goto common;

	case ODIV:
	case OLDIV:
		ccom(l);
		if(vconst(l) == 0 &amp;&amp; !side(r)) {
			*n = *l;
			break;
		}
		ccom(r);
		t = vconst(r);
		if(t == 0) {
			diag(n, &#34;divide check&#34;);
			*n = *r;
			break;
		}
		if(t == 1) {
			*n = *l;
			break;
		}
		goto common;

	case OSUB:
		ccom(r);
		if(r-&gt;op == OCONST) {
			if(typefd[r-&gt;type-&gt;etype]) {
				n-&gt;op = OADD;
				r-&gt;fconst = -r-&gt;fconst;
				goto loop;
			} else {
				n-&gt;op = OADD;
				r-&gt;vconst = -r-&gt;vconst;
				goto loop;
			}
		}
		ccom(l);
		goto common;

	case OXOR:
	case OOR:
	case OADD:
		ccom(l);
		if(vconst(l) == 0) {
			*n = *r;
			goto loop;
		}
		ccom(r);
		if(vconst(r) == 0) {
			*n = *l;
			break;
		}
		goto commute;

	case OAND:
		ccom(l);
		ccom(r);
		if(vconst(l) == 0 &amp;&amp; !side(r)) {
			*n = *l;
			break;
		}
		if(vconst(r) == 0 &amp;&amp; !side(l)) {
			*n = *r;
			break;
		}

	commute:
		/* look for commutative constant */
		if(r-&gt;op == OCONST) {
			if(l-&gt;op == n-&gt;op) {
				if(l-&gt;left-&gt;op == OCONST) {
					n-&gt;right = l-&gt;right;
					l-&gt;right = r;
					goto loop;
				}
				if(l-&gt;right-&gt;op == OCONST) {
					n-&gt;right = l-&gt;left;
					l-&gt;left = r;
					goto loop;
				}
			}
		}
		if(l-&gt;op == OCONST) {
			if(r-&gt;op == n-&gt;op) {
				if(r-&gt;left-&gt;op == OCONST) {
					n-&gt;left = r-&gt;right;
					r-&gt;right = l;
					goto loop;
				}
				if(r-&gt;right-&gt;op == OCONST) {
					n-&gt;left = r-&gt;left;
					r-&gt;left = l;
					goto loop;
				}
			}
		}
		goto common;

	case OANDAND:
		ccom(l);
		if(vconst(l) == 0) {
			*n = *l;
			break;
		}
		ccom(r);
		goto common;

	case OOROR:
		ccom(l);
		if(l-&gt;op == OCONST &amp;&amp; l-&gt;vconst != 0) {
			*n = *l;
			n-&gt;vconst = 1;
			break;
		}
		ccom(r);
		goto common;

	default:
		if(l != Z)
			ccom(l);
		if(r != Z)
			ccom(r);
	common:
		if(l != Z)
		if(l-&gt;op != OCONST)
			break;
		if(r != Z)
		if(r-&gt;op != OCONST)
			break;
		evconst(n);
	}
}

/*	OEQ, ONE, OLE, OLS, OLT, OLO, OGE, OHS, OGT, OHI */
static char *cmps[12] = 
{
	&#34;==&#34;, &#34;!=&#34;, &#34;&lt;=&#34;, &#34;&lt;=&#34;, &#34;&lt;&#34;, &#34;&lt;&#34;, &#34;&gt;=&#34;, &#34;&gt;=&#34;, &#34;&gt;&#34;, &#34;&gt;&#34;,
};

/* 128-bit numbers */
typedef struct Big Big;
struct Big
{
	vlong a;
	uvlong b;
};
static int
cmp(Big x, Big y)
{
	if(x.a != y.a){
		if(x.a &lt; y.a)
			return -1;
		return 1;
	}
	if(x.b != y.b){
		if(x.b &lt; y.b)
			return -1;
		return 1;
	}
	return 0;
}
static Big
add(Big x, int y)
{
	uvlong ob;
	
	ob = x.b;
	x.b += y;
	if(y &gt; 0 &amp;&amp; x.b &lt; ob)
		x.a++;
	if(y &lt; 0 &amp;&amp; x.b &gt; ob)
		x.a--;
	return x;
} 

Big
big(vlong a, uvlong b)
{
	Big x;

	x.a = a;
	x.b = b;
	return x;
}

int
compar(Node *n, int reverse)
{
	Big lo, hi, x;
	int op;
	char xbuf[40], cmpbuf[50];
	Node *l, *r;
	Type *lt, *rt;

	/*
	 * The point of this function is to diagnose comparisons 
	 * that can never be true or that look misleading because
	 * of the `usual arithmetic conversions&#39;.  As an example 
	 * of the latter, if x is a ulong, then if(x &lt;= -1) really means
	 * if(x &lt;= 0xFFFFFFFF), while if(x &lt;= -1LL) really means
	 * what it says (but 8c compiles it wrong anyway).
	 */

	if(reverse){
		r = n-&gt;left;
		l = n-&gt;right;
		op = comrel[relindex(n-&gt;op)];
	}else{
		l = n-&gt;left;
		r = n-&gt;right;
		op = n-&gt;op;
	}

	/*
	 * Skip over left casts to find out the original expression range.
	 */
	while(l-&gt;op == OCAST)
		l = l-&gt;left;
	if(l-&gt;op == OCONST)
		return 0;
	lt = l-&gt;type;
	if(l-&gt;op == ONAME &amp;&amp; l-&gt;sym-&gt;type){
		lt = l-&gt;sym-&gt;type;
		if(lt-&gt;etype == TARRAY)
			lt = lt-&gt;link;
	}
	if(lt == T)
		return 0;
	if(lt-&gt;etype == TXXX || lt-&gt;etype &gt; TUVLONG)
		return 0;
	
	/*
	 * Skip over the right casts to find the on-screen value.
	 */
	if(r-&gt;op != OCONST)
		return 0;
	while(r-&gt;oldop == OCAST &amp;&amp; !r-&gt;xcast)
		r = r-&gt;left;
	rt = r-&gt;type;
	if(rt == T)
		return 0;

	x.b = r-&gt;vconst;
	x.a = 0;
	if((rt-&gt;etype&amp;1) &amp;&amp; r-&gt;vconst &lt; 0)	/* signed negative */
		x.a = ~0ULL;

	if((lt-&gt;etype&amp;1)==0){
		/* unsigned */
		lo = big(0, 0);
		if(lt-&gt;width == 8)
			hi = big(0, ~0ULL);
		else
			hi = big(0, (1LL&lt;&lt;(l-&gt;type-&gt;width*8))-1);
	}else{
		lo = big(~0ULL, -(1LL&lt;&lt;(l-&gt;type-&gt;width*8-1)));
		hi = big(0, (1LL&lt;&lt;(l-&gt;type-&gt;width*8-1))-1);
	}

	switch(op){
	case OLT:
	case OLO:
	case OGE:
	case OHS:
		if(cmp(x, lo) &lt;= 0)
			goto useless;
		if(cmp(x, add(hi, 1)) &gt;= 0)
			goto useless;
		break;
	case OLE:
	case OLS:
	case OGT:
	case OHI:
		if(cmp(x, add(lo, -1)) &lt;= 0)
			goto useless;
		if(cmp(x, hi) &gt;= 0)
			goto useless;
		break;
	case OEQ:
	case ONE:
		/*
		 * Don&#39;t warn about comparisons if the expression
		 * is as wide as the value: the compiler-supplied casts
		 * will make both outcomes possible.
		 */
		if(lt-&gt;width &gt;= rt-&gt;width &amp;&amp; debug[&#39;w&#39;] &lt; 2)
			return 0;
		if(cmp(x, lo) &lt; 0 || cmp(x, hi) &gt; 0)
			goto useless;
		break;
	}
	return 0;

useless:
	if((x.a==0 &amp;&amp; x.b&lt;=9) || (x.a==~0LL &amp;&amp; x.b &gt;= -9ULL))
		snprint(xbuf, sizeof xbuf, &#34;%lld&#34;, x.b);
	else if(x.a == 0)
		snprint(xbuf, sizeof xbuf, &#34;%#llux&#34;, x.b);
	else
		snprint(xbuf, sizeof xbuf, &#34;%#llx&#34;, x.b);
	if(reverse)
		snprint(cmpbuf, sizeof cmpbuf, &#34;%s %s %T&#34;,
			xbuf, cmps[relindex(n-&gt;op)], lt);
	else
		snprint(cmpbuf, sizeof cmpbuf, &#34;%T %s %s&#34;,
			lt, cmps[relindex(n-&gt;op)], xbuf);
	warn(n, &#34;useless or misleading comparison: %s&#34;, cmpbuf);
	return 0;
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

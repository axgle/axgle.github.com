<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/const.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/const.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;
#define	TUP(x,y)	(((x)&lt;&lt;16)|(y))

static Val toflt(Val);
static Val toint(Val);
static Val tostr(Val);
static void overflow(Val, Type*);
static Val copyval(Val);

/*
 * truncate float literal fv to 32-bit or 64-bit precision
 * according to type; return truncated value.
 */
Mpflt*
truncfltlit(Mpflt *oldv, Type *t)
{
	double d;
	float f;
	Mpflt *fv;

	if(t == T)
		return oldv;

	fv = mal(sizeof *fv);
	*fv = *oldv;

	// convert large precision literal floating
	// into limited precision (float64 or float32)
	// botch -- this assumes that compiler fp
	//    has same precision as runtime fp
	switch(t-&gt;etype) {
	case TFLOAT64:
		d = mpgetflt(fv);
		mpmovecflt(fv, d);
		break;

	case TFLOAT32:
		d = mpgetflt(fv);
		f = d;
		d = f;
		mpmovecflt(fv, d);
		break;
	}
	return fv;
}

/*
 * convert n, if literal, to type t.
 * implicit conversion.
 */
void
convlit(Node **np, Type *t)
{
	return convlit1(np, t, 0);
}

/*
 * convert n, if literal, to type t.
 * return a new node if necessary
 * (if n is a named constant, can&#39;t edit n-&gt;type directly).
 */
void
convlit1(Node **np, Type *t, int explicit)
{
	int ct, et;
	Node *n, *nn;

	n = *np;
	if(n == N || t == T || n-&gt;type == T || isideal(t) || n-&gt;type == t)
		return;
	if(!explicit &amp;&amp; !isideal(n-&gt;type))
		return;

//dump(&#34;convlit1&#34;, n);
	if(n-&gt;op == OLITERAL) {
		nn = nod(OXXX, N, N);
		*nn = *n;
		n = nn;
		*np = n;
	}
//dump(&#34;convlit2&#34;, n);

	switch(n-&gt;op) {
	default:
		if(n-&gt;type-&gt;etype == TIDEAL) {
			convlit(&amp;n-&gt;left, t);
			convlit(&amp;n-&gt;right, t);
			n-&gt;type = t;
		}
		return;
	case OLITERAL:
		break;
	case OLSH:
	case ORSH:
		convlit1(&amp;n-&gt;left, t, explicit);
		t = n-&gt;left-&gt;type;
		if(t != T &amp;&amp; !isint[t-&gt;etype]) {
			yyerror(&#34;invalid operation: %#N (shift of type %T)&#34;, n, t);
			t = T;
		}
		n-&gt;type = t;
		return;
	}
	// avoided repeated calculations, errors
	if(cvttype(n-&gt;type, t) == 1) {
		n-&gt;type = t;
		return;
	}

	ct = consttype(n);
	if(ct &lt; 0)
		goto bad;

	et = t-&gt;etype;
	if(et == TINTER) {
		if(ct == CTNIL &amp;&amp; n-&gt;type == types[TNIL]) {
			n-&gt;type = t;
			return;
		}
		defaultlit(np, T);
		return;
	}

	switch(ct) {
	default:
		goto bad;

	case CTNIL:
		switch(et) {
		default:
			yyerror(&#34;cannot use nil as %T&#34;, t);
			n-&gt;type = T;
			goto bad;

		case TSTRING:
			// let normal conversion code handle it
			return;

		case TARRAY:
			if(!isslice(t))
				goto bad;
			break;

		case TPTR32:
		case TPTR64:
		case TINTER:
		case TMAP:
		case TCHAN:
		case TFUNC:
			break;
		}
		break;

	case CTSTR:
	case CTBOOL:
		if(et != n-&gt;type-&gt;etype)
			goto bad;
		break;

	case CTINT:
	case CTFLT:
		ct = n-&gt;val.ctype;
		if(isint[et]) {
			if(ct == CTFLT)
				n-&gt;val = toint(n-&gt;val);
			else if(ct != CTINT)
				goto bad;
			overflow(n-&gt;val, t);
		} else if(isfloat[et]) {
			if(ct == CTINT)
				n-&gt;val = toflt(n-&gt;val);
			else if(ct != CTFLT)
				goto bad;
			overflow(n-&gt;val, t);
			n-&gt;val.u.fval = truncfltlit(n-&gt;val.u.fval, t);
		} else if(et == TSTRING &amp;&amp; ct == CTINT &amp;&amp; explicit)
			n-&gt;val = tostr(n-&gt;val);
		else
			goto bad;
	}
	n-&gt;type = t;
	return;

bad:
	if(isideal(n-&gt;type)) {
		defaultlit(&amp;n, T);
		*np = n;
	}
	return;
}

static Val
copyval(Val v)
{
	Mpint *i;
	Mpflt *f;

	switch(v.ctype) {
	case CTINT:
		i = mal(sizeof(*i));
		mpmovefixfix(i, v.u.xval);
		v.u.xval = i;
		break;
	case CTFLT:
		f = mal(sizeof(*f));
		mpmovefltflt(f, v.u.fval);
		v.u.fval = f;
		break;
	}
	return v;
}

static Val
toflt(Val v)
{
	Mpflt *f;

	if(v.ctype == CTINT) {
		f = mal(sizeof(*f));
		mpmovefixflt(f, v.u.xval);
		v.ctype = CTFLT;
		v.u.fval = f;
	}
	return v;
}

static Val
toint(Val v)
{
	Mpint *i;

	if(v.ctype == CTFLT) {
		i = mal(sizeof(*i));
		if(mpmovefltfix(i, v.u.fval) &lt; 0)
			yyerror(&#34;constant %#F truncated to integer&#34;, v.u.fval);
		v.ctype = CTINT;
		v.u.xval = i;
	}
	return v;
}

static void
overflow(Val v, Type *t)
{
	// v has already been converted
	// to appropriate form for t.
	if(t == T || t-&gt;etype == TIDEAL)
		return;
	switch(v.ctype) {
	case CTINT:
		if(!isint[t-&gt;etype])
			fatal(&#34;overflow: %T integer constant&#34;, t);
		if(mpcmpfixfix(v.u.xval, minintval[t-&gt;etype]) &lt; 0
		|| mpcmpfixfix(v.u.xval, maxintval[t-&gt;etype]) &gt; 0)
			yyerror(&#34;constant %B overflows %T&#34;, v.u.xval, t);
		break;
	case CTFLT:
		if(!isfloat[t-&gt;etype])
			fatal(&#34;overflow: %T floating-point constant&#34;, t);
		if(mpcmpfltflt(v.u.fval, minfltval[t-&gt;etype]) &lt; 0
		|| mpcmpfltflt(v.u.fval, maxfltval[t-&gt;etype]) &gt; 0)
			yyerror(&#34;constant %#F overflows %T&#34;, v.u.fval, t);
		break;
	}
}

static Val
tostr(Val v)
{
	Rune rune;
	int l;
	Strlit *s;

	switch(v.ctype) {
	case CTINT:
		if(mpcmpfixfix(v.u.xval, minintval[TINT]) &lt; 0
		|| mpcmpfixfix(v.u.xval, maxintval[TINT]) &gt; 0)
			yyerror(&#34;overflow in int -&gt; string&#34;);
		rune = mpgetfix(v.u.xval);
		l = runelen(rune);
		s = mal(sizeof(*s)+l);
		s-&gt;len = l;
		runetochar((char*)s-&gt;s, &amp;rune);
		memset(&amp;v, 0, sizeof v);
		v.ctype = CTSTR;
		v.u.sval = s;
		break;

	case CTFLT:
		yyerror(&#34;no float -&gt; string&#34;);

	case CTNIL:
		memset(&amp;v, 0, sizeof v);
		v.ctype = CTSTR;
		v.u.sval = mal(sizeof *s);
		break;
	}
	return v;
}

int
consttype(Node *n)
{
	if(n == N || n-&gt;op != OLITERAL)
		return -1;
	return n-&gt;val.ctype;
}

int
isconst(Node *n, int ct)
{
	return consttype(n) == ct;
}

/*
 * if n is constant, rewrite as OLITERAL node.
 */
void
evconst(Node *n)
{
	Node *nl, *nr;
	int32 len;
	Strlit *str;
	int wl, wr, lno, et;
	Val v;
	Mpint b;

	switch(n-&gt;op) {
	case OMAKE:
	case OMAKEMAP:
	case OMAKESLICE:
	case OMAKECHAN:
	case ODCLCONST:
		return;
	}

	nl = n-&gt;left;
	if(nl == N || nl-&gt;type == T)
		return;
	if(consttype(nl) &lt; 0)
		return;
	wl = nl-&gt;type-&gt;etype;
	if(isint[wl] || isfloat[wl])
		wl = TIDEAL;

	nr = n-&gt;right;
	if(nr == N)
		goto unary;
	if(nr-&gt;type == T)
		return;
	if(consttype(nr) &lt; 0)
		return;
	wr = nr-&gt;type-&gt;etype;
	if(isint[wr] || isfloat[wr])
		wr = TIDEAL;

	// check for compatible general types (numeric, string, etc)
	if(wl != wr)
		goto illegal;

	// check for compatible types.
	switch(n-&gt;op) {
	default:
		// ideal const mixes with anything but otherwise must match.
		if(nl-&gt;type-&gt;etype != TIDEAL) {
			defaultlit(&amp;nr, nl-&gt;type);
			n-&gt;right = nr;
		}
		if(nr-&gt;type-&gt;etype != TIDEAL) {
			defaultlit(&amp;nl, nr-&gt;type);
			n-&gt;left = nl;
		}
		if(nl-&gt;type-&gt;etype != nr-&gt;type-&gt;etype)
			goto illegal;
		break;

	case OLSH:
	case ORSH:
		// right must be unsigned.
		// left can be ideal.
		defaultlit(&amp;nr, types[TUINT]);
		n-&gt;right = nr;
		if(nr-&gt;type &amp;&amp; (issigned[nr-&gt;type-&gt;etype] || !isint[nr-&gt;type-&gt;etype]))
			goto illegal;
		break;
	}

	// copy numeric value to avoid modifying
	// n-&gt;left, in case someone still refers to it (e.g. iota).
	v = nl-&gt;val;
	if(wl == TIDEAL)
		v = copyval(v);

	// since wl == wr,
	// the only way v.ctype != nr-&gt;val.ctype
	// is when one is CTINT and the other CTFLT.
	// make both CTFLT.
	if(v.ctype != nr-&gt;val.ctype) {
		v = toflt(v);
		nr-&gt;val = toflt(nr-&gt;val);
	}

	// run op
	switch(TUP(n-&gt;op, v.ctype)) {
	default:
	illegal:
		if(!n-&gt;diag) {
			yyerror(&#34;illegal constant expression: %T %O %T&#34;,
				nl-&gt;type, n-&gt;op, nr-&gt;type);
			n-&gt;diag = 1;
		}
		return;

	case TUP(OADD, CTINT):
		mpaddfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(OSUB, CTINT):
		mpsubfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(OMUL, CTINT):
		mpmulfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(ODIV, CTINT):
		if(mpcmpfixc(nr-&gt;val.u.xval, 0) == 0) {
			yyerror(&#34;division by zero&#34;);
			mpmovecfix(v.u.xval, 1);
			break;
		}
		mpdivfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(OMOD, CTINT):
		if(mpcmpfixc(nr-&gt;val.u.xval, 0) == 0) {
			yyerror(&#34;division by zero&#34;);
			mpmovecfix(v.u.xval, 1);
			break;
		}
		mpmodfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;

	case TUP(OLSH, CTINT):
		mplshfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(ORSH, CTINT):
		mprshfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(OOR, CTINT):
		mporfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(OAND, CTINT):
		mpandfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(OANDNOT, CTINT):
		mpandnotfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(OXOR, CTINT):
		mpxorfixfix(v.u.xval, nr-&gt;val.u.xval);
		break;
	case TUP(OADD, CTFLT):
		mpaddfltflt(v.u.fval, nr-&gt;val.u.fval);
		break;
	case TUP(OSUB, CTFLT):
		mpsubfltflt(v.u.fval, nr-&gt;val.u.fval);
		break;
	case TUP(OMUL, CTFLT):
		mpmulfltflt(v.u.fval, nr-&gt;val.u.fval);
		break;
	case TUP(ODIV, CTFLT):
		if(mpcmpfltc(nr-&gt;val.u.fval, 0) == 0) {
			yyerror(&#34;division by zero&#34;);
			mpmovecflt(v.u.fval, 1.0);
			break;
		}
		mpdivfltflt(v.u.fval, nr-&gt;val.u.fval);
		break;

	case TUP(OEQ, CTNIL):
		goto settrue;
	case TUP(ONE, CTNIL):
		goto setfalse;

	case TUP(OEQ, CTINT):
		if(mpcmpfixfix(v.u.xval, nr-&gt;val.u.xval) == 0)
			goto settrue;
		goto setfalse;
	case TUP(ONE, CTINT):
		if(mpcmpfixfix(v.u.xval, nr-&gt;val.u.xval) != 0)
			goto settrue;
		goto setfalse;
	case TUP(OLT, CTINT):
		if(mpcmpfixfix(v.u.xval, nr-&gt;val.u.xval) &lt; 0)
			goto settrue;
		goto setfalse;
	case TUP(OLE, CTINT):
		if(mpcmpfixfix(v.u.xval, nr-&gt;val.u.xval) &lt;= 0)
			goto settrue;
		goto setfalse;
	case TUP(OGE, CTINT):
		if(mpcmpfixfix(v.u.xval, nr-&gt;val.u.xval) &gt;= 0)
			goto settrue;
		goto setfalse;
	case TUP(OGT, CTINT):
		if(mpcmpfixfix(v.u.xval, nr-&gt;val.u.xval) &gt; 0)
			goto settrue;
		goto setfalse;

	case TUP(OEQ, CTFLT):
		if(mpcmpfltflt(v.u.fval, nr-&gt;val.u.fval) == 0)
			goto settrue;
		goto setfalse;
	case TUP(ONE, CTFLT):
		if(mpcmpfltflt(v.u.fval, nr-&gt;val.u.fval) != 0)
			goto settrue;
		goto setfalse;
	case TUP(OLT, CTFLT):
		if(mpcmpfltflt(v.u.fval, nr-&gt;val.u.fval) &lt; 0)
			goto settrue;
		goto setfalse;
	case TUP(OLE, CTFLT):
		if(mpcmpfltflt(v.u.fval, nr-&gt;val.u.fval) &lt;= 0)
			goto settrue;
		goto setfalse;
	case TUP(OGE, CTFLT):
		if(mpcmpfltflt(v.u.fval, nr-&gt;val.u.fval) &gt;= 0)
			goto settrue;
		goto setfalse;
	case TUP(OGT, CTFLT):
		if(mpcmpfltflt(v.u.fval, nr-&gt;val.u.fval) &gt; 0)
			goto settrue;
		goto setfalse;

	case TUP(OEQ, CTSTR):
		if(cmpslit(nl, nr) == 0)
			goto settrue;
		goto setfalse;
	case TUP(ONE, CTSTR):
		if(cmpslit(nl, nr) != 0)
			goto settrue;
		goto setfalse;
	case TUP(OLT, CTSTR):
		if(cmpslit(nl, nr) &lt; 0)
			goto settrue;
		goto setfalse;
	case TUP(OLE, CTSTR):
		if(cmpslit(nl, nr) &lt;= 0)
			goto settrue;
		goto setfalse;
	case TUP(OGE, CTSTR):
		if(cmpslit(nl, nr) &gt;= 0l)
			goto settrue;
		goto setfalse;
	case TUP(OGT, CTSTR):
		if(cmpslit(nl, nr) &gt; 0)
			goto settrue;
		goto setfalse;
	case TUP(OADDSTR, CTSTR):
		len = v.u.sval-&gt;len + nr-&gt;val.u.sval-&gt;len;
		str = mal(sizeof(*str) + len);
		str-&gt;len = len;
		memcpy(str-&gt;s, v.u.sval-&gt;s, v.u.sval-&gt;len);
		memcpy(str-&gt;s+v.u.sval-&gt;len, nr-&gt;val.u.sval-&gt;s, nr-&gt;val.u.sval-&gt;len);
		str-&gt;len = len;
		v.u.sval = str;
		break;

	case TUP(OOROR, CTBOOL):
		if(v.u.bval || nr-&gt;val.u.bval)
			goto settrue;
		goto setfalse;
	case TUP(OANDAND, CTBOOL):
		if(v.u.bval &amp;&amp; nr-&gt;val.u.bval)
			goto settrue;
		goto setfalse;
	case TUP(OEQ, CTBOOL):
		if(v.u.bval == nr-&gt;val.u.bval)
			goto settrue;
		goto setfalse;
	case TUP(ONE, CTBOOL):
		if(v.u.bval != nr-&gt;val.u.bval)
			goto settrue;
		goto setfalse;
	}
	goto ret;

unary:
	// copy numeric value to avoid modifying
	// nl, in case someone still refers to it (e.g. iota).
	v = nl-&gt;val;
	if(wl == TIDEAL)
		v = copyval(v);

	switch(TUP(n-&gt;op, v.ctype)) {
	default:
		if(!n-&gt;diag) {
			yyerror(&#34;illegal constant expression %O %T&#34;, n-&gt;op, nl-&gt;type);
			n-&gt;diag = 1;
		}
		return;

	case TUP(OCONV, CTNIL):
	case TUP(OARRAYBYTESTR, CTNIL):
		if(n-&gt;type-&gt;etype == TSTRING) {
			v = tostr(v);
			nl-&gt;type = n-&gt;type;
			break;
		}
		// fall through
	case TUP(OCONV, CTINT):
	case TUP(OCONV, CTFLT):
	case TUP(OCONV, CTSTR):
		convlit1(&amp;nl, n-&gt;type, 1);
		break;

	case TUP(OPLUS, CTINT):
		break;
	case TUP(OMINUS, CTINT):
		mpnegfix(v.u.xval);
		break;
	case TUP(OCOM, CTINT):
		et = Txxx;
		if(nl-&gt;type != T)
			et = nl-&gt;type-&gt;etype;

		// calculate the mask in b
		// result will be (a ^ mask)
		switch(et) {
		default:
			// signed guys change sign
			mpmovecfix(&amp;b, -1);
			break;

		case TUINT8:
		case TUINT16:
		case TUINT32:
		case TUINT64:
		case TUINT:
		case TUINTPTR:
			// unsigned guys invert their bits
			mpmovefixfix(&amp;b, maxintval[et]);
			break;
		}
		mpxorfixfix(v.u.xval, &amp;b);
		break;

	case TUP(OPLUS, CTFLT):
		break;
	case TUP(OMINUS, CTFLT):
		mpnegflt(v.u.fval);
		break;

	case TUP(ONOT, CTBOOL):
		if(!v.u.bval)
			goto settrue;
		goto setfalse;
	}

ret:
	// rewrite n in place.
	*n = *nl;
	n-&gt;val = v;

	// check range.
	lno = setlineno(n);
	overflow(v, n-&gt;type);
	lineno = lno;

	// truncate precision for non-ideal float.
	if(v.ctype == CTFLT &amp;&amp; n-&gt;type-&gt;etype != TIDEAL)
		n-&gt;val.u.fval = truncfltlit(v.u.fval, n-&gt;type);
	return;

settrue:
	*n = *nodbool(1);
	return;

setfalse:
	*n = *nodbool(0);
	return;
}

Node*
nodlit(Val v)
{
	Node *n;

	n = nod(OLITERAL, N, N);
	n-&gt;val = v;
	switch(v.ctype) {
	default:
		fatal(&#34;nodlit ctype %d&#34;, v.ctype);
	case CTSTR:
		n-&gt;type = idealstring;
		break;
	case CTBOOL:
		n-&gt;type = idealbool;
		break;
	case CTINT:
	case CTFLT:
		n-&gt;type = types[TIDEAL];
		break;
	case CTNIL:
		n-&gt;type = types[TNIL];
		break;
	}
	return n;
}

// TODO(rsc): combine with convlit
void
defaultlit(Node **np, Type *t)
{
	int lno;
	Node *n, *nn;

	n = *np;
	if(n == N || !isideal(n-&gt;type))
		return;

	switch(n-&gt;op) {
	case OLITERAL:
		nn = nod(OXXX, N, N);
		*nn = *n;
		n = nn;
		*np = n;
		break;
	case OLSH:
	case ORSH:
		defaultlit(&amp;n-&gt;left, t);
		t = n-&gt;left-&gt;type;
		if(t != T &amp;&amp; !isint[t-&gt;etype]) {
			yyerror(&#34;invalid operation: %#N (shift of type %T)&#34;, n, t);
			t = T;
		}
		n-&gt;type = t;
		return;
	default:
		if(n-&gt;left == N) {
			dump(&#34;defaultlit&#34;, n);
			fatal(&#34;defaultlit&#34;);
		}
		defaultlit(&amp;n-&gt;left, t);
		defaultlit(&amp;n-&gt;right, t);
		if(n-&gt;type == idealbool || n-&gt;type == idealstring)
			n-&gt;type = types[n-&gt;type-&gt;etype];
		else
			n-&gt;type = n-&gt;left-&gt;type;
		return;
	}

	lno = setlineno(n);
	switch(n-&gt;val.ctype) {
	default:
		if(t != T) {
			convlit(np, t);
			break;
		}
		if(n-&gt;val.ctype == CTNIL) {
			lineno = lno;
			yyerror(&#34;use of untyped nil&#34;);
			n-&gt;type = T;
			break;
		}
		if(n-&gt;val.ctype == CTSTR) {
			n-&gt;type = types[TSTRING];
			break;
		}
		yyerror(&#34;defaultlit: unknown literal: %#N&#34;, n);
		break;
	case CTBOOL:
		n-&gt;type = types[TBOOL];
		break;
	case CTINT:
		n-&gt;type = types[TINT];
		if(t != T) {
			if(isint[t-&gt;etype])
				n-&gt;type = t;
			else if(isfloat[t-&gt;etype]) {
				n-&gt;type = t;
				n-&gt;val = toflt(n-&gt;val);
			}
		}
		overflow(n-&gt;val, n-&gt;type);
		break;
	case CTFLT:
		n-&gt;type = types[TFLOAT];
		if(t != T) {
			if(isfloat[t-&gt;etype])
				n-&gt;type = t;
			else if(isint[t-&gt;etype]) {
				n-&gt;type = t;
				n-&gt;val = toint(n-&gt;val);
			}
		}
		overflow(n-&gt;val, n-&gt;type);
		break;
	}
	lineno = lno;
}

/*
 * defaultlit on both nodes simultaneously;
 * if they&#39;re both ideal going in they better
 * get the same type going out.
 */
void
defaultlit2(Node **lp, Node **rp, int force)
{
	Node *l, *r;

	l = *lp;
	r = *rp;
	if(l-&gt;type == T || r-&gt;type == T)
		return;
	if(!isideal(l-&gt;type)) {
		convlit(rp, l-&gt;type);
		return;
	}
	if(!isideal(r-&gt;type)) {
		convlit(lp, r-&gt;type);
		return;
	}
	if(!force)
		return;
	if(isconst(l, CTFLT) || isconst(r, CTFLT)) {
		convlit(lp, types[TFLOAT]);
		convlit(rp, types[TFLOAT]);
		return;
	}
	convlit(lp, types[TINT]);
	convlit(rp, types[TINT]);
}

int
cmpslit(Node *l, Node *r)
{
	int32 l1, l2, i, m;
	char *s1, *s2;

	l1 = l-&gt;val.u.sval-&gt;len;
	l2 = r-&gt;val.u.sval-&gt;len;
	s1 = l-&gt;val.u.sval-&gt;s;
	s2 = r-&gt;val.u.sval-&gt;s;

	m = l1;
	if(l2 &lt; m)
		m = l2;

	for(i=0; i&lt;m; i++) {
		if(s1[i] == s2[i])
			continue;
		if(s1[i] &gt; s2[i])
			return +1;
		return -1;
	}
	if(l1 == l2)
		return 0;
	if(l1 &gt; l2)
		return +1;
	return -1;
}

int
smallintconst(Node *n)
{
	if(n-&gt;op == OLITERAL &amp;&amp; n-&gt;type != T)
	switch(simtype[n-&gt;type-&gt;etype]) {
	case TINT8:
	case TUINT8:
	case TINT16:
	case TUINT16:
	case TINT32:
	case TUINT32:
	case TBOOL:
	case TPTR32:
		return 1;
	}
	return 0;
}

long
nonnegconst(Node *n)
{
	if(n-&gt;op == OLITERAL &amp;&amp; n-&gt;type != T)
	switch(simtype[n-&gt;type-&gt;etype]) {
	case TINT8:
	case TUINT8:
	case TINT16:
	case TUINT16:
	case TINT32:
	case TUINT32:
	case TINT64:
	case TUINT64:
	case TIDEAL:
		// check negative and 2^31
		if(mpcmpfixfix(n-&gt;val.u.xval, minintval[TUINT32]) &lt; 0
		|| mpcmpfixfix(n-&gt;val.u.xval, maxintval[TINT32]) &gt; 0)
			break;
		return mpgetfix(n-&gt;val.u.xval);
	}
	return -1;
}

/*
 * convert x to type et and back to int64
 * for sign extension and truncation.
 */
int64
iconv(int64 x, int et)
{
	switch(et) {
	case TINT8:
		x = (int8)x;
		break;
	case TUINT8:
		x = (uint8)x;
		break;
	case TINT16:
		x = (int16)x;
		break;
	case TUINT16:
		x = (uint64)x;
		break;
	case TINT32:
		x = (int32)x;
		break;
	case TUINT32:
		x = (uint32)x;
		break;
	case TINT64:
	case TUINT64:
		break;
	}
	return x;
}

/*
 * convert constant val to type t; leave in con.
 * for back end.
 */
void
convconst(Node *con, Type *t, Val *val)
{
	int64 i;
	int tt;

	tt = simsimtype(t);

	// copy the constant for conversion
	nodconst(con, types[TINT8], 0);
	con-&gt;type = t;
	con-&gt;val = *val;

	if(isint[tt]) {
		con-&gt;val.ctype = CTINT;
		con-&gt;val.u.xval = mal(sizeof *con-&gt;val.u.xval);
		switch(val-&gt;ctype) {
		default:
			fatal(&#34;convconst ctype=%d %lT&#34;, val-&gt;ctype, t);
		case CTINT:
			i = mpgetfix(val-&gt;u.xval);
			break;
		case CTBOOL:
			i = val-&gt;u.bval;
			break;
		case CTNIL:
			i = 0;
			break;
		}
		i = iconv(i, tt);
		mpmovecfix(con-&gt;val.u.xval, i);
		return;
	}

	if(isfloat[tt]) {
		if(con-&gt;val.ctype == CTINT) {
			con-&gt;val.ctype = CTFLT;
			con-&gt;val.u.fval = mal(sizeof *con-&gt;val.u.fval);
			mpmovefixflt(con-&gt;val.u.fval, val-&gt;u.xval);
		}
		if(con-&gt;val.ctype != CTFLT)
			fatal(&#34;convconst ctype=%d %T&#34;, con-&gt;val.ctype, t);
		if(!isfloat[tt]) {
			// easy to handle, but can it happen?
			fatal(&#34;convconst CTINT %T&#34;, t);
		}
		if(tt == TFLOAT32)
			con-&gt;val.u.fval = truncfltlit(con-&gt;val.u.fval, t);
		return;
	}

	fatal(&#34;convconst %lT constant&#34;, t);

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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/funct.c</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/funct.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/funct.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/funct.c
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

#include	&#34;cc.h&#34;

typedef	struct	Ftab	Ftab;
struct	Ftab
{
	char	op;
	char*	name;
	char	typ;
};
typedef	struct	Gtab	Gtab;
struct	Gtab
{
	char	etype;
	char*	name;
};

Ftab	ftabinit[OEND];
Gtab	gtabinit[NTYPE];

int
isfunct(Node *n)
{
	Type *t, *t1;
	Funct *f;
	Node *l;
	Sym *s;
	int o;

	o = n-&gt;op;
	if(n-&gt;left == Z)
		goto no;
	t = n-&gt;left-&gt;type;
	if(t == T)
		goto no;
	f = t-&gt;funct;

	switch(o) {
	case OAS:	// put cast on rhs
	case OASI:
	case OASADD:
	case OASAND:
	case OASASHL:
	case OASASHR:
	case OASDIV:
	case OASLDIV:
	case OASLMOD:
	case OASLMUL:
	case OASLSHR:
	case OASMOD:
	case OASMUL:
	case OASOR:
	case OASSUB:
	case OASXOR:
		if(n-&gt;right == Z)
			goto no;
		t1 = n-&gt;right-&gt;type;
		if(t1 == T)
			goto no;
		if(t1-&gt;funct == f)
			break;

		l = new(OXXX, Z, Z);
		*l = *n-&gt;right;

		n-&gt;right-&gt;left = l;
		n-&gt;right-&gt;right = Z;
		n-&gt;right-&gt;type = t;
		n-&gt;right-&gt;op = OCAST;

		if(!isfunct(n-&gt;right))
			prtree(n, &#34;isfunc !&#34;);
		break;

	case OCAST:	// t f(T) or T f(t)
		t1 = n-&gt;type;
		if(t1 == T)
			goto no;
		if(f != nil) {
			s = f-&gt;castfr[t1-&gt;etype];
			if(s == S)
				goto no;
			n-&gt;right = n-&gt;left;
			goto build;
		}
		f = t1-&gt;funct;
		if(f != nil) {
			s = f-&gt;castto[t-&gt;etype];
			if(s == S)
				goto no;
			n-&gt;right = n-&gt;left;
			goto build;
		}
		goto no;
	}

	if(f == nil)
		goto no;
	s = f-&gt;sym[o];
	if(s == S)
		goto no;

	/*
	 * the answer is yes,
	 * now we rewrite the node
	 * and give diagnostics
	 */
	switch(o) {
	default:
		diag(n, &#34;isfunct op missing %O\n&#34;, o);
		goto bad;

	case OADD:	// T f(T, T)
	case OAND:
	case OASHL:
	case OASHR:
	case ODIV:
	case OLDIV:
	case OLMOD:
	case OLMUL:
	case OLSHR:
	case OMOD:
	case OMUL:
	case OOR:
	case OSUB:
	case OXOR:

	case OEQ:	// int f(T, T)
	case OGE:
	case OGT:
	case OHI:
	case OHS:
	case OLE:
	case OLO:
	case OLS:
	case OLT:
	case ONE:
		if(n-&gt;right == Z)
			goto bad;
		t1 = n-&gt;right-&gt;type;
		if(t1 == T)
			goto bad;
		if(t1-&gt;funct != f)
			goto bad;
		n-&gt;right = new(OLIST, n-&gt;left, n-&gt;right);
		break;

	case OAS:	// structure copies done by the compiler
	case OASI:
		goto no;

	case OASADD:	// T f(T*, T)
	case OASAND:
	case OASASHL:
	case OASASHR:
	case OASDIV:
	case OASLDIV:
	case OASLMOD:
	case OASLMUL:
	case OASLSHR:
	case OASMOD:
	case OASMUL:
	case OASOR:
	case OASSUB:
	case OASXOR:
		if(n-&gt;right == Z)
			goto bad;
		t1 = n-&gt;right-&gt;type;
		if(t1 == T)
			goto bad;
		if(t1-&gt;funct != f)
			goto bad;
		n-&gt;right = new(OLIST, new(OADDR, n-&gt;left, Z), n-&gt;right);
		break;

	case OPOS:	// T f(T)
	case ONEG:
	case ONOT:
	case OCOM:
		n-&gt;right = n-&gt;left;
		break;


	}

build:
	l = new(ONAME, Z, Z);
	l-&gt;sym = s;
	l-&gt;type = s-&gt;type;
	l-&gt;etype = s-&gt;type-&gt;etype;
	l-&gt;xoffset = s-&gt;offset;
	l-&gt;class = s-&gt;class;
	tcomo(l, 0);

	n-&gt;op = OFUNC;
	n-&gt;left = l;
	n-&gt;type = l-&gt;type-&gt;link;
	if(tcompat(n, T, l-&gt;type, tfunct))
		goto bad;
	if(tcoma(n-&gt;left, n-&gt;right, l-&gt;type-&gt;down, 1))
		goto bad;
	return 1;

no:
	return 0;

bad:
	diag(n, &#34;cant rewrite typestr for op %O\n&#34;, o);
	prtree(n, &#34;isfunct&#34;);
	n-&gt;type = T;
	return 1;
}

void
dclfunct(Type *t, Sym *s)
{
	Funct *f;
	Node *n;
	Type *f1, *f2, *f3, *f4;
	int o, i, c;
	char str[100];

	if(t-&gt;funct)
		return;

	// recognize generated tag of dorm _%d_
	if(t-&gt;tag == S)
		goto bad;
	for(i=0; c = t-&gt;tag-&gt;name[i]; i++) {
		if(c == &#39;_&#39;) {
			if(i == 0 || t-&gt;tag-&gt;name[i+1] == 0)
				continue;
			break;
		}
		if(c &lt; &#39;0&#39; || c &gt; &#39;9&#39;)
			break;
	}
	if(c == 0)
		goto bad;

	f = alloc(sizeof(*f));
	for(o=0; o&lt;sizeof(f-&gt;sym); o++)
		f-&gt;sym[o] = S;

	t-&gt;funct = f;

	f1 = typ(TFUNC, t);
	f1-&gt;down = copytyp(t);
	f1-&gt;down-&gt;down = t;

	f2 = typ(TFUNC, types[TINT]);
	f2-&gt;down = copytyp(t);
	f2-&gt;down-&gt;down = t;

	f3 = typ(TFUNC, t);
	f3-&gt;down = typ(TIND, t);
	f3-&gt;down-&gt;down = t;

	f4 = typ(TFUNC, t);
	f4-&gt;down = t;

	for(i=0;; i++) {
		o = ftabinit[i].op;
		if(o == OXXX)
			break;
		sprint(str, &#34;%s_%s_&#34;, t-&gt;tag-&gt;name, ftabinit[i].name);
		n = new(ONAME, Z, Z);
		n-&gt;sym = slookup(str);
		f-&gt;sym[o] = n-&gt;sym;
		switch(ftabinit[i].typ) {
		default:
			diag(Z, &#34;dclfunct op missing %d\n&#34;, ftabinit[i].typ);
			break;

		case 1:	// T f(T,T)	+
			dodecl(xdecl, CEXTERN, f1, n);
			break;

		case 2:	// int f(T,T)	==
			dodecl(xdecl, CEXTERN, f2, n);
			break;

		case 3:	// void f(T*,T)	+=
			dodecl(xdecl, CEXTERN, f3, n);
			break;

		case 4:	// T f(T)	~
			dodecl(xdecl, CEXTERN, f4, n);
			break;
		}
	}
	for(i=0;; i++) {
		o = gtabinit[i].etype;
		if(o == TXXX)
			break;

		/*
		 * OCAST types T1 _T2_T1_(T2)
		 */
		sprint(str, &#34;_%s%s_&#34;, gtabinit[i].name, t-&gt;tag-&gt;name);
		n = new(ONAME, Z, Z);
		n-&gt;sym = slookup(str);
		f-&gt;castto[o] = n-&gt;sym;

		f1 = typ(TFUNC, t);
		f1-&gt;down = types[o];
		dodecl(xdecl, CEXTERN, f1, n);

		sprint(str, &#34;%s_%s_&#34;, t-&gt;tag-&gt;name, gtabinit[i].name);
		n = new(ONAME, Z, Z);
		n-&gt;sym = slookup(str);
		f-&gt;castfr[o] = n-&gt;sym;

		f1 = typ(TFUNC, types[o]);
		f1-&gt;down = t;
		dodecl(xdecl, CEXTERN, f1, n);
	}
	return;
bad:
	diag(Z, &#34;dclfunct bad %T %s\n&#34;, t, s-&gt;name);
}

Gtab	gtabinit[NTYPE] =
{
	TCHAR,		&#34;c&#34;,
	TUCHAR,		&#34;uc&#34;,
	TSHORT,		&#34;h&#34;,
	TUSHORT,	&#34;uh&#34;,
	TINT,		&#34;i&#34;,
	TUINT,		&#34;ui&#34;,
	TLONG,		&#34;l&#34;,
	TULONG,		&#34;ul&#34;,
	TVLONG,		&#34;v&#34;,
	TUVLONG,	&#34;uv&#34;,
	TFLOAT,		&#34;f&#34;,
	TDOUBLE,	&#34;d&#34;,
	TXXX
};

Ftab	ftabinit[OEND] =
{
	OADD,		&#34;add&#34;,		1,
	OAND,		&#34;and&#34;,		1,
	OASHL,		&#34;ashl&#34;,		1,
	OASHR,		&#34;ashr&#34;,		1,
	ODIV,		&#34;div&#34;,		1,
	OLDIV,		&#34;ldiv&#34;,		1,
	OLMOD,		&#34;lmod&#34;,		1,
	OLMUL,		&#34;lmul&#34;,		1,
	OLSHR,		&#34;lshr&#34;,		1,
	OMOD,		&#34;mod&#34;,		1,
	OMUL,		&#34;mul&#34;,		1,
	OOR,		&#34;or&#34;,		1,
	OSUB,		&#34;sub&#34;,		1,
	OXOR,		&#34;xor&#34;,		1,

	OEQ,		&#34;eq&#34;,		2,
	OGE,		&#34;ge&#34;,		2,
	OGT,		&#34;gt&#34;,		2,
	OHI,		&#34;hi&#34;,		2,
	OHS,		&#34;hs&#34;,		2,
	OLE,		&#34;le&#34;,		2,
	OLO,		&#34;lo&#34;,		2,
	OLS,		&#34;ls&#34;,		2,
	OLT,		&#34;lt&#34;,		2,
	ONE,		&#34;ne&#34;,		2,

	OASADD,		&#34;asadd&#34;,	3,
	OASAND,		&#34;asand&#34;,	3,
	OASASHL,	&#34;asashl&#34;,	3,
	OASASHR,	&#34;asashr&#34;,	3,
	OASDIV,		&#34;asdiv&#34;,	3,
	OASLDIV,	&#34;asldiv&#34;,	3,
	OASLMOD,	&#34;aslmod&#34;,	3,
	OASLMUL,	&#34;aslmul&#34;,	3,
	OASLSHR,	&#34;aslshr&#34;,	3,
	OASMOD,		&#34;asmod&#34;,	3,
	OASMUL,		&#34;asmul&#34;,	3,
	OASOR,		&#34;asor&#34;,		3,
	OASSUB,		&#34;assub&#34;,	3,
	OASXOR,		&#34;asxor&#34;,	3,

	OPOS,		&#34;pos&#34;,		4,
	ONEG,		&#34;neg&#34;,		4,
	OCOM,		&#34;com&#34;,		4,
	ONOT,		&#34;not&#34;,		4,

//	OPOSTDEC,
//	OPOSTINC,
//	OPREDEC,
//	OPREINC,

	OXXX,
};

//	Node*	nodtestv;

//	Node*	nodvpp;
//	Node*	nodppv;
//	Node*	nodvmm;
//	Node*	nodmmv;
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

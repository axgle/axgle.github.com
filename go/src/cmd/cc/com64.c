<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/com64.c</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/com64.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/com64.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/com64.c
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

/*
 * this is machine depend, but it is totally
 * common on all of the 64-bit symulating machines.
 */

#define	FNX	100	/* botch -- redefinition */

Node*	nodaddv;
Node*	nodsubv;
Node*	nodmulv;
Node*	noddivv;
Node*	noddivvu;
Node*	nodmodv;
Node*	nodmodvu;
Node*	nodlshv;
Node*	nodrshav;
Node*	nodrshlv;
Node*	nodandv;
Node*	nodorv;
Node*	nodxorv;
Node*	nodnegv;
Node*	nodcomv;

Node*	nodtestv;
Node*	nodeqv;
Node*	nodnev;
Node*	nodlev;
Node*	nodltv;
Node*	nodgev;
Node*	nodgtv;
Node*	nodhiv;
Node*	nodhsv;
Node*	nodlov;
Node*	nodlsv;

Node*	nodf2v;
Node*	nodd2v;
Node*	nodp2v;
Node*	nodsi2v;
Node*	nodui2v;
Node*	nodsl2v;
Node*	nodul2v;
Node*	nodsh2v;
Node*	noduh2v;
Node*	nodsc2v;
Node*	noduc2v;

Node*	nodv2f;
Node*	nodv2d;
Node*	nodv2ui;
Node*	nodv2si;
Node*	nodv2ul;
Node*	nodv2sl;
Node*	nodv2uh;
Node*	nodv2sh;
Node*	nodv2uc;
Node*	nodv2sc;

Node*	nodvpp;
Node*	nodppv;
Node*	nodvmm;
Node*	nodmmv;

Node*	nodvasop;

char	etconv[NTYPE];	/* for _vasop */
Init	initetconv[] =
{
	TCHAR,		1,	0,
	TUCHAR,		2,	0,
	TSHORT,		3,	0,
	TUSHORT,	4,	0,
	TLONG,		5,	0,
	TULONG,		6,	0,
	TVLONG,		7,	0,
	TUVLONG,	8,	0,
	TINT,		9,	0,
	TUINT,		10,	0,
	-1,		0,	0,
};

Node*
fvn(char *name, int type)
{
	Node *n;

	n = new(ONAME, Z, Z);
	n-&gt;sym = slookup(name);
	n-&gt;sym-&gt;sig = SIGINTERN;
	if(fntypes[type] == 0)
		fntypes[type] = typ(TFUNC, types[type]);
	n-&gt;type = fntypes[type];
	n-&gt;etype = type;
	n-&gt;class = CGLOBL;
	n-&gt;addable = 10;
	n-&gt;complex = 0;
	return n;
}

void
com64init(void)
{
	Init *p;

	nodaddv = fvn(&#34;_addv&#34;, TVLONG);
	nodsubv = fvn(&#34;_subv&#34;, TVLONG);
	nodmulv = fvn(&#34;_mulv&#34;, TVLONG);
	noddivv = fvn(&#34;_divv&#34;, TVLONG);
	noddivvu = fvn(&#34;_divvu&#34;, TVLONG);
	nodmodv = fvn(&#34;_modv&#34;, TVLONG);
	nodmodvu = fvn(&#34;_modvu&#34;, TVLONG);
	nodlshv = fvn(&#34;_lshv&#34;, TVLONG);
	nodrshav = fvn(&#34;_rshav&#34;, TVLONG);
	nodrshlv = fvn(&#34;_rshlv&#34;, TVLONG);
	nodandv = fvn(&#34;_andv&#34;, TVLONG);
	nodorv = fvn(&#34;_orv&#34;, TVLONG);
	nodxorv = fvn(&#34;_xorv&#34;, TVLONG);
	nodnegv = fvn(&#34;_negv&#34;, TVLONG);
	nodcomv = fvn(&#34;_comv&#34;, TVLONG);

	nodtestv = fvn(&#34;_testv&#34;, TLONG);
	nodeqv = fvn(&#34;_eqv&#34;, TLONG);
	nodnev = fvn(&#34;_nev&#34;, TLONG);
	nodlev = fvn(&#34;_lev&#34;, TLONG);
	nodltv = fvn(&#34;_ltv&#34;, TLONG);
	nodgev = fvn(&#34;_gev&#34;, TLONG);
	nodgtv = fvn(&#34;_gtv&#34;, TLONG);
	nodhiv = fvn(&#34;_hiv&#34;, TLONG);
	nodhsv = fvn(&#34;_hsv&#34;, TLONG);
	nodlov = fvn(&#34;_lov&#34;, TLONG);
	nodlsv = fvn(&#34;_lsv&#34;, TLONG);

	nodf2v = fvn(&#34;_f2v&#34;, TVLONG);
	nodd2v = fvn(&#34;_d2v&#34;, TVLONG);
	nodp2v = fvn(&#34;_p2v&#34;, TVLONG);
	nodsi2v = fvn(&#34;_si2v&#34;, TVLONG);
	nodui2v = fvn(&#34;_ui2v&#34;, TVLONG);
	nodsl2v = fvn(&#34;_sl2v&#34;, TVLONG);
	nodul2v = fvn(&#34;_ul2v&#34;, TVLONG);
	nodsh2v = fvn(&#34;_sh2v&#34;, TVLONG);
	noduh2v = fvn(&#34;_uh2v&#34;, TVLONG);
	nodsc2v = fvn(&#34;_sc2v&#34;, TVLONG);
	noduc2v = fvn(&#34;_uc2v&#34;, TVLONG);

	nodv2f = fvn(&#34;_v2f&#34;, TFLOAT);
	nodv2d = fvn(&#34;_v2d&#34;, TDOUBLE);
	nodv2sl = fvn(&#34;_v2sl&#34;, TLONG);
	nodv2ul = fvn(&#34;_v2ul&#34;, TULONG);
	nodv2si = fvn(&#34;_v2si&#34;, TINT);
	nodv2ui = fvn(&#34;_v2ui&#34;, TUINT);
	nodv2sh = fvn(&#34;_v2sh&#34;, TSHORT);
	nodv2uh = fvn(&#34;_v2ul&#34;, TUSHORT);
	nodv2sc = fvn(&#34;_v2sc&#34;, TCHAR);
	nodv2uc = fvn(&#34;_v2uc&#34;, TUCHAR);

	nodvpp = fvn(&#34;_vpp&#34;, TVLONG);
	nodppv = fvn(&#34;_ppv&#34;, TVLONG);
	nodvmm = fvn(&#34;_vmm&#34;, TVLONG);
	nodmmv = fvn(&#34;_mmv&#34;, TVLONG);

	nodvasop = fvn(&#34;_vasop&#34;, TVLONG);

	for(p = initetconv; p-&gt;code &gt;= 0; p++)
		etconv[p-&gt;code] = p-&gt;value;
}

int
com64(Node *n)
{
	Node *l, *r, *a, *t;
	int lv, rv;

	if(n-&gt;type == 0)
		return 0;

	l = n-&gt;left;
	r = n-&gt;right;

	lv = 0;
	if(l &amp;&amp; l-&gt;type &amp;&amp; typev[l-&gt;type-&gt;etype])
		lv = 1;
	rv = 0;
	if(r &amp;&amp; r-&gt;type &amp;&amp; typev[r-&gt;type-&gt;etype])
		rv = 1;

	if(lv) {
		switch(n-&gt;op) {
		case OEQ:
			a = nodeqv;
			goto setbool;
		case ONE:
			a = nodnev;
			goto setbool;
		case OLE:
			a = nodlev;
			goto setbool;
		case OLT:
			a = nodltv;
			goto setbool;
		case OGE:
			a = nodgev;
			goto setbool;
		case OGT:
			a = nodgtv;
			goto setbool;
		case OHI:
			a = nodhiv;
			goto setbool;
		case OHS:
			a = nodhsv;
			goto setbool;
		case OLO:
			a = nodlov;
			goto setbool;
		case OLS:
			a = nodlsv;
			goto setbool;

		case OANDAND:
		case OOROR:
			if(machcap(n))
				return 1;

			if(rv) {
				r = new(OFUNC, nodtestv, r);
				n-&gt;right = r;
				r-&gt;complex = FNX;
				r-&gt;op = OFUNC;
				r-&gt;type = types[TLONG];
			}

		case OCOND:
		case ONOT:
			if(machcap(n))
				return 1;

			l = new(OFUNC, nodtestv, l);
			n-&gt;left = l;
			l-&gt;complex = FNX;
			l-&gt;op = OFUNC;
			l-&gt;type = types[TLONG];
			n-&gt;complex = FNX;
			return 1;
		}
	}

	if(rv) {
		if(machcap(n))
			return 1;
		switch(n-&gt;op) {
		case OANDAND:
		case OOROR:
			r = new(OFUNC, nodtestv, r);
			n-&gt;right = r;
			r-&gt;complex = FNX;
			r-&gt;op = OFUNC;
			r-&gt;type = types[TLONG];
			return 1;
		}
	}

	if(typev[n-&gt;type-&gt;etype]) {
		if(machcap(n))
			return 1;
		switch(n-&gt;op) {
		default:
			diag(n, &#34;unknown vlong %O&#34;, n-&gt;op);
		case OFUNC:
			n-&gt;complex = FNX;
		case ORETURN:
		case OAS:
		case OIND:
			return 1;
		case OADD:
			a = nodaddv;
			goto setbop;
		case OSUB:
			a = nodsubv;
			goto setbop;
		case OMUL:
		case OLMUL:
			a = nodmulv;
			goto setbop;
		case ODIV:
			a = noddivv;
			goto setbop;
		case OLDIV:
			a = noddivvu;
			goto setbop;
		case OMOD:
			a = nodmodv;
			goto setbop;
		case OLMOD:
			a = nodmodvu;
			goto setbop;
		case OASHL:
			a = nodlshv;
			goto setbop;
		case OASHR:
			a = nodrshav;
			goto setbop;
		case OLSHR:
			a = nodrshlv;
			goto setbop;
		case OAND:
			a = nodandv;
			goto setbop;
		case OOR:
			a = nodorv;
			goto setbop;
		case OXOR:
			a = nodxorv;
			goto setbop;
		case OPOSTINC:
			a = nodvpp;
			goto setvinc;
		case OPOSTDEC:
			a = nodvmm;
			goto setvinc;
		case OPREINC:
			a = nodppv;
			goto setvinc;
		case OPREDEC:
			a = nodmmv;
			goto setvinc;
		case ONEG:
			a = nodnegv;
			goto setfnx;
		case OCOM:
			a = nodcomv;
			goto setfnx;
		case OCAST:
			switch(l-&gt;type-&gt;etype) {
			case TCHAR:
				a = nodsc2v;
				goto setfnxl;
			case TUCHAR:
				a = noduc2v;
				goto setfnxl;
			case TSHORT:
				a = nodsh2v;
				goto setfnxl;
			case TUSHORT:
				a = noduh2v;
				goto setfnxl;
			case TINT:
				a = nodsi2v;
				goto setfnx;
			case TUINT:
				a = nodui2v;
				goto setfnx;
			case TLONG:
				a = nodsl2v;
				goto setfnx;
			case TULONG:
				a = nodul2v;
				goto setfnx;
			case TFLOAT:
				a = nodf2v;
				goto setfnx;
			case TDOUBLE:
				a = nodd2v;
				goto setfnx;
			case TIND:
				a = nodp2v;
				goto setfnx;
			}
			diag(n, &#34;unknown %T-&gt;vlong cast&#34;, l-&gt;type);
			return 1;
		case OASADD:
			a = nodaddv;
			goto setasop;
		case OASSUB:
			a = nodsubv;
			goto setasop;
		case OASMUL:
		case OASLMUL:
			a = nodmulv;
			goto setasop;
		case OASDIV:
			a = noddivv;
			goto setasop;
		case OASLDIV:
			a = noddivvu;
			goto setasop;
		case OASMOD:
			a = nodmodv;
			goto setasop;
		case OASLMOD:
			a = nodmodvu;
			goto setasop;
		case OASASHL:
			a = nodlshv;
			goto setasop;
		case OASASHR:
			a = nodrshav;
			goto setasop;
		case OASLSHR:
			a = nodrshlv;
			goto setasop;
		case OASAND:
			a = nodandv;
			goto setasop;
		case OASOR:
			a = nodorv;
			goto setasop;
		case OASXOR:
			a = nodxorv;
			goto setasop;
		}
	}

	if(typefd[n-&gt;type-&gt;etype] &amp;&amp; l &amp;&amp; l-&gt;op == OFUNC) {
		switch(n-&gt;op) {
		case OASADD:
		case OASSUB:
		case OASMUL:
		case OASLMUL:
		case OASDIV:
		case OASLDIV:
		case OASMOD:
		case OASLMOD:
		case OASASHL:
		case OASASHR:
		case OASLSHR:
		case OASAND:
		case OASOR:
		case OASXOR:
			if(l-&gt;right &amp;&amp; typev[l-&gt;right-&gt;etype]) {
				diag(n, &#34;sorry float &lt;asop&gt; vlong not implemented\n&#34;);
			}
		}
	}

	if(n-&gt;op == OCAST) {
		if(l-&gt;type &amp;&amp; typev[l-&gt;type-&gt;etype]) {
			if(machcap(n))
				return 1;
			switch(n-&gt;type-&gt;etype) {
			case TDOUBLE:
				a = nodv2d;
				goto setfnx;
			case TFLOAT:
				a = nodv2f;
				goto setfnx;
			case TLONG:
				a = nodv2sl;
				goto setfnx;
			case TULONG:
				a = nodv2ul;
				goto setfnx;
			case TINT:
				a = nodv2si;
				goto setfnx;
			case TUINT:
				a = nodv2ui;
				goto setfnx;
			case TSHORT:
				a = nodv2sh;
				goto setfnx;
			case TUSHORT:
				a = nodv2uh;
				goto setfnx;
			case TCHAR:
				a = nodv2sc;
				goto setfnx;
			case TUCHAR:
				a = nodv2uc;
				goto setfnx;
			case TIND:	// small pun here
				a = nodv2ul;
				goto setfnx;
			}
			diag(n, &#34;unknown vlong-&gt;%T cast&#34;, n-&gt;type);
			return 1;
		}
	}

	return 0;

setbop:
	n-&gt;left = a;
	n-&gt;right = new(OLIST, l, r);
	n-&gt;complex = FNX;
	n-&gt;op = OFUNC;
	return 1;

setfnxl:
	l = new(OCAST, l, 0);
	l-&gt;type = types[TLONG];
	l-&gt;complex = l-&gt;left-&gt;complex;

setfnx:
	n-&gt;left = a;
	n-&gt;right = l;
	n-&gt;complex = FNX;
	n-&gt;op = OFUNC;
	return 1;

setvinc:
	n-&gt;left = a;
	l = new(OADDR, l, Z);
	l-&gt;type = typ(TIND, l-&gt;left-&gt;type);
	n-&gt;right = new(OLIST, l, r);
	n-&gt;complex = FNX;
	n-&gt;op = OFUNC;
	return 1;

setbool:
	if(machcap(n))
		return 1;
	n-&gt;left = a;
	n-&gt;right = new(OLIST, l, r);
	n-&gt;complex = FNX;
	n-&gt;op = OFUNC;
	n-&gt;type = types[TLONG];
	return 1;

setasop:
	if(l-&gt;op == OFUNC) {
		l = l-&gt;right;
		goto setasop;
	}

	t = new(OCONST, 0, 0);
	t-&gt;vconst = etconv[l-&gt;type-&gt;etype];
	t-&gt;type = types[TLONG];
	t-&gt;addable = 20;
	r = new(OLIST, t, r);

	t = new(OADDR, a, 0);
	t-&gt;type = typ(TIND, a-&gt;type);
	r = new(OLIST, t, r);

	t = new(OADDR, l, 0);
	t-&gt;type = typ(TIND, l-&gt;type);
	r = new(OLIST, t, r);

	n-&gt;left = nodvasop;
	n-&gt;right = r;
	n-&gt;complex = FNX;
	n-&gt;op = OFUNC;

	return 1;
}

void
bool64(Node *n)
{
	Node *n1;

	if(machcap(Z))
		return;
	if(typev[n-&gt;type-&gt;etype]) {
		n1 = new(OXXX, 0, 0);
		*n1 = *n;

		n-&gt;right = n1;
		n-&gt;left = nodtestv;
		n-&gt;complex = FNX;
		n-&gt;addable = 0;
		n-&gt;op = OFUNC;
		n-&gt;type = types[TLONG];
	}
}

/*
 * more machine depend stuff.
 * this is common for 8,16,32,64 bit machines.
 * this is common for ieee machines.
 */
double
convvtof(vlong v)
{
	double d;

	d = v;		/* BOTCH */
	return d;
}

vlong
convftov(double d)
{
	vlong v;


	v = d;		/* BOTCH */
	return v;
}

double
convftox(double d, int et)
{

	if(!typefd[et])
		diag(Z, &#34;bad type in castftox %s&#34;, tnames[et]);
	return d;
}

vlong
convvtox(vlong c, int et)
{
	int n;

	n = 8 * ewidth[et];
	c &amp;= MASK(n);
	if(!typeu[et])
		if(c &amp; SIGN(n))
			c |= ~MASK(n);
	return c;
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

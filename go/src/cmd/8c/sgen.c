<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8c/sgen.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8c/sgen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8c/sgen.c
// http://code.google.com/p/inferno-os/source/browse/utils/8c/sgen.c
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

Prog*
gtext(Sym *s, int32 stkoff)
{
	gpseudo(ATEXT, s, nodconst(stkoff));
	p-&gt;to.type = D_CONST2;
	p-&gt;to.offset2 = argsize();
	return p;
}

void
noretval(int n)
{

	if(n &amp; 1) {
		gins(ANOP, Z, Z);
		p-&gt;to.type = REGRET;
	}
	if(n &amp; 2) {
		gins(ANOP, Z, Z);
		p-&gt;to.type = FREGRET;
	}
}

/* welcome to commute */
static void
commute(Node *n)
{
	Node *l, *r;

	l = n-&gt;left;
	r = n-&gt;right;
	if(r-&gt;complex &gt; l-&gt;complex) {
		n-&gt;left = r;
		n-&gt;right = l;
	}
}

void
indexshift(Node *n)
{
	int g;

	if(!typechlp[n-&gt;type-&gt;etype])
		return;
	simplifyshift(n);
	if(n-&gt;op == OASHL &amp;&amp; n-&gt;right-&gt;op == OCONST){
		g = vconst(n-&gt;right);
		if(g &gt;= 0 &amp;&amp; g &lt; 4)
			n-&gt;addable = 7;
	}
}

/*
 *	calculate addressability as follows
 *		NAME ==&gt; 10/11		name+value(SB/SP)
 *		REGISTER ==&gt; 12		register
 *		CONST ==&gt; 20		$value
 *		*(20) ==&gt; 21		value
 *		&amp;(10) ==&gt; 13		$name+value(SB)
 *		&amp;(11) ==&gt; 1		$name+value(SP)
 *		(13) + (20) ==&gt; 13	fold constants
 *		(1) + (20) ==&gt; 1	fold constants
 *		*(13) ==&gt; 10		back to name
 *		*(1) ==&gt; 11		back to name
 *
 *		(20) * (X) ==&gt; 7	multiplier in indexing
 *		(X,7) + (13,1) ==&gt; 8	adder in indexing (addresses)
 *		(8) ==&gt; &amp;9(OINDEX)	index, almost addressable
 *
 *	calculate complexity (number of registers)
 */
void
xcom(Node *n)
{
	Node *l, *r;
	int g;

	if(n == Z)
		return;
	l = n-&gt;left;
	r = n-&gt;right;
	n-&gt;complex = 0;
	n-&gt;addable = 0;
	switch(n-&gt;op) {
	case OCONST:
		n-&gt;addable = 20;
		break;

	case ONAME:
		n-&gt;addable = 10;
		if(n-&gt;class == CPARAM || n-&gt;class == CAUTO)
			n-&gt;addable = 11;
		break;

	case OEXREG:
		n-&gt;addable = 10;
		break;

	case OREGISTER:
		n-&gt;addable = 12;
		break;

	case OINDREG:
		n-&gt;addable = 12;
		break;

	case OADDR:
		xcom(l);
		if(l-&gt;addable == 10)
			n-&gt;addable = 13;
		else
		if(l-&gt;addable == 11)
			n-&gt;addable = 1;
		break;

	case OADD:
		xcom(l);
		xcom(r);
		if(n-&gt;type-&gt;etype != TIND)
			break;

		switch(r-&gt;addable) {
		case 20:
			switch(l-&gt;addable) {
			case 1:
			case 13:
			commadd:
				l-&gt;type = n-&gt;type;
				*n = *l;
				l = new(0, Z, Z);
				*l = *(n-&gt;left);
				l-&gt;xoffset += r-&gt;vconst;
				n-&gt;left = l;
				r = n-&gt;right;
				goto brk;
			}
			break;

		case 1:
		case 13:
		case 10:
		case 11:
			/* l is the base, r is the index */
			if(l-&gt;addable != 20)
				n-&gt;addable = 8;
			break;
		}
		switch(l-&gt;addable) {
		case 20:
			switch(r-&gt;addable) {
			case 13:
			case 1:
				r = n-&gt;left;
				l = n-&gt;right;
				n-&gt;left = l;
				n-&gt;right = r;
				goto commadd;
			}
			break;

		case 13:
		case 1:
		case 10:
		case 11:
			/* r is the base, l is the index */
			if(r-&gt;addable != 20)
				n-&gt;addable = 8;
			break;
		}
		if(n-&gt;addable == 8 &amp;&amp; !side(n)) {
			indx(n);
			l = new1(OINDEX, idx.basetree, idx.regtree);
			l-&gt;scale = idx.scale;
			l-&gt;addable = 9;
			l-&gt;complex = l-&gt;right-&gt;complex;
			l-&gt;type = l-&gt;left-&gt;type;
			n-&gt;op = OADDR;
			n-&gt;left = l;
			n-&gt;right = Z;
			n-&gt;addable = 8;
			break;
		}
		break;

	case OINDEX:
		xcom(l);
		xcom(r);
		n-&gt;addable = 9;
		break;

	case OIND:
		xcom(l);
		if(l-&gt;op == OADDR) {
			l = l-&gt;left;
			l-&gt;type = n-&gt;type;
			*n = *l;
			return;
		}
		switch(l-&gt;addable) {
		case 20:
			n-&gt;addable = 21;
			break;
		case 1:
			n-&gt;addable = 11;
			break;
		case 13:
			n-&gt;addable = 10;
			break;
		}
		break;

	case OASHL:
		xcom(l);
		xcom(r);
		indexshift(n);
		break;

	case OMUL:
	case OLMUL:
		xcom(l);
		xcom(r);
		g = vlog(l);
		if(g &gt;= 0) {
			n-&gt;left = r;
			n-&gt;right = l;
			l = r;
			r = n-&gt;right;
		}
		g = vlog(r);
		if(g &gt;= 0) {
			n-&gt;op = OASHL;
			r-&gt;vconst = g;
			r-&gt;type = types[TINT];
			indexshift(n);
			break;
		}
commute(n);
		break;

	case OASLDIV:
		xcom(l);
		xcom(r);
		g = vlog(r);
		if(g &gt;= 0) {
			n-&gt;op = OASLSHR;
			r-&gt;vconst = g;
			r-&gt;type = types[TINT];
		}
		break;

	case OLDIV:
		xcom(l);
		xcom(r);
		g = vlog(r);
		if(g &gt;= 0) {
			n-&gt;op = OLSHR;
			r-&gt;vconst = g;
			r-&gt;type = types[TINT];
			indexshift(n);
			break;
		}
		break;

	case OASLMOD:
		xcom(l);
		xcom(r);
		g = vlog(r);
		if(g &gt;= 0) {
			n-&gt;op = OASAND;
			r-&gt;vconst--;
		}
		break;

	case OLMOD:
		xcom(l);
		xcom(r);
		g = vlog(r);
		if(g &gt;= 0) {
			n-&gt;op = OAND;
			r-&gt;vconst--;
		}
		break;

	case OASMUL:
	case OASLMUL:
		xcom(l);
		xcom(r);
		g = vlog(r);
		if(g &gt;= 0) {
			n-&gt;op = OASASHL;
			r-&gt;vconst = g;
		}
		break;

	case OLSHR:
	case OASHR:
		xcom(l);
		xcom(r);
		indexshift(n);
		break;

	default:
		if(l != Z)
			xcom(l);
		if(r != Z)
			xcom(r);
		break;
	}
brk:
	if(n-&gt;addable &gt;= 10)
		return;
	if(l != Z)
		n-&gt;complex = l-&gt;complex;
	if(r != Z) {
		if(r-&gt;complex == n-&gt;complex)
			n-&gt;complex = r-&gt;complex+1;
		else
		if(r-&gt;complex &gt; n-&gt;complex)
			n-&gt;complex = r-&gt;complex;
	}
	if(n-&gt;complex == 0)
		n-&gt;complex++;

	if(com64(n))
		return;

	switch(n-&gt;op) {

	case OFUNC:
		n-&gt;complex = FNX;
		break;

	case OLMOD:
	case OMOD:
	case OLMUL:
	case OLDIV:
	case OMUL:
	case ODIV:
	case OASLMUL:
	case OASLDIV:
	case OASLMOD:
	case OASMUL:
	case OASDIV:
	case OASMOD:
		if(r-&gt;complex &gt;= l-&gt;complex) {
			n-&gt;complex = l-&gt;complex + 3;
			if(r-&gt;complex &gt; n-&gt;complex)
				n-&gt;complex = r-&gt;complex;
		} else {
			n-&gt;complex = r-&gt;complex + 3;
			if(l-&gt;complex &gt; n-&gt;complex)
				n-&gt;complex = l-&gt;complex;
		}
		break;

	case OLSHR:
	case OASHL:
	case OASHR:
	case OASLSHR:
	case OASASHL:
	case OASASHR:
		if(r-&gt;complex &gt;= l-&gt;complex) {
			n-&gt;complex = l-&gt;complex + 2;
			if(r-&gt;complex &gt; n-&gt;complex)
				n-&gt;complex = r-&gt;complex;
		} else {
			n-&gt;complex = r-&gt;complex + 2;
			if(l-&gt;complex &gt; n-&gt;complex)
				n-&gt;complex = l-&gt;complex;
		}
		break;

	case OADD:
	case OXOR:
	case OAND:
	case OOR:
		/*
		 * immediate operators, make const on right
		 */
		if(l-&gt;op == OCONST) {
			n-&gt;left = r;
			n-&gt;right = l;
		}
		break;

	case OEQ:
	case ONE:
	case OLE:
	case OLT:
	case OGE:
	case OGT:
	case OHI:
	case OHS:
	case OLO:
	case OLS:
		/*
		 * compare operators, make const on left
		 */
		if(r-&gt;op == OCONST) {
			n-&gt;left = r;
			n-&gt;right = l;
			n-&gt;op = invrel[relindex(n-&gt;op)];
		}
		break;
	}
}

void
indx(Node *n)
{
	Node *l, *r;

	if(debug[&#39;x&#39;])
		prtree(n, &#34;indx&#34;);

	l = n-&gt;left;
	r = n-&gt;right;
	if(l-&gt;addable == 1 || l-&gt;addable == 13 || r-&gt;complex &gt; l-&gt;complex) {
		n-&gt;right = l;
		n-&gt;left = r;
		l = r;
		r = n-&gt;right;
	}
	if(l-&gt;addable != 7) {
		idx.regtree = l;
		idx.scale = 1;
	} else
	if(l-&gt;right-&gt;addable == 20) {
		idx.regtree = l-&gt;left;
		idx.scale = 1 &lt;&lt; l-&gt;right-&gt;vconst;
	} else
	if(l-&gt;left-&gt;addable == 20) {
		idx.regtree = l-&gt;right;
		idx.scale = 1 &lt;&lt; l-&gt;left-&gt;vconst;
	} else
		diag(n, &#34;bad index&#34;);

	idx.basetree = r;
	if(debug[&#39;x&#39;]) {
		print(&#34;scale = %d\n&#34;, idx.scale);
		prtree(idx.regtree, &#34;index&#34;);
		prtree(idx.basetree, &#34;base&#34;);
	}
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

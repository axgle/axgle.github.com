<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/scon.c</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/scon.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/scon.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/scon.c
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

static Node*
acast(Type *t, Node *n)
{
	if(n-&gt;type-&gt;etype != t-&gt;etype || n-&gt;op == OBIT) {
		n = new1(OCAST, n, Z);
		if(nocast(n-&gt;left-&gt;type, t))
			*n = *n-&gt;left;
		n-&gt;type = t;
	}
	return n;
}


void
evconst(Node *n)
{
	Node *l, *r;
	int et, isf;
	vlong v;
	double d;

	if(n == Z || n-&gt;type == T)
		return;

	et = n-&gt;type-&gt;etype;
	isf = typefd[et];

	l = n-&gt;left;
	r = n-&gt;right;

	d = 0;
	v = 0;

	switch(n-&gt;op) {
	default:
		return;

	case ONEG:
		if(isf)
			d = -l-&gt;fconst;
		else
			v = -l-&gt;vconst;
		break;

	case OCOM:
		v = ~l-&gt;vconst;
		break;

	case OCAST:
		if(et == TVOID)
			return;
		et = l-&gt;type-&gt;etype;
		if(isf) {
			if(typefd[et])
				d = l-&gt;fconst;
			else
				d = l-&gt;vconst;
		} else {
			if(typefd[et])
				v = l-&gt;fconst;
			else
				v = convvtox(l-&gt;vconst, n-&gt;type-&gt;etype);
		}
		break;

	case OCONST:
		break;

	case OADD:
		if(isf)
			d = l-&gt;fconst + r-&gt;fconst;
		else {
			v = l-&gt;vconst + r-&gt;vconst;
		}
		break;

	case OSUB:
		if(isf)
			d = l-&gt;fconst - r-&gt;fconst;
		else
			v = l-&gt;vconst - r-&gt;vconst;
		break;

	case OMUL:
		if(isf)
			d = l-&gt;fconst * r-&gt;fconst;
		else {
			v = l-&gt;vconst * r-&gt;vconst;
		}
		break;

	case OLMUL:
		v = (uvlong)l-&gt;vconst * (uvlong)r-&gt;vconst;
		break;


	case ODIV:
		if(vconst(r) == 0) {
			warn(n, &#34;divide by zero&#34;);
			return;
		}
		if(isf)
			d = l-&gt;fconst / r-&gt;fconst;
		else
			v = l-&gt;vconst / r-&gt;vconst;
		break;

	case OLDIV:
		if(vconst(r) == 0) {
			warn(n, &#34;divide by zero&#34;);
			return;
		}
		v = (uvlong)l-&gt;vconst / (uvlong)r-&gt;vconst;
		break;

	case OMOD:
		if(vconst(r) == 0) {
			warn(n, &#34;modulo by zero&#34;);
			return;
		}
		v = l-&gt;vconst % r-&gt;vconst;
		break;

	case OLMOD:
		if(vconst(r) == 0) {
			warn(n, &#34;modulo by zero&#34;);
			return;
		}
		v = (uvlong)l-&gt;vconst % (uvlong)r-&gt;vconst;
		break;

	case OAND:
		v = l-&gt;vconst &amp; r-&gt;vconst;
		break;

	case OOR:
		v = l-&gt;vconst | r-&gt;vconst;
		break;

	case OXOR:
		v = l-&gt;vconst ^ r-&gt;vconst;
		break;

	case OLSHR:
		v = (uvlong)l-&gt;vconst &gt;&gt; r-&gt;vconst;
		break;

	case OASHR:
		v = l-&gt;vconst &gt;&gt; r-&gt;vconst;
		break;

	case OASHL:
		v = l-&gt;vconst &lt;&lt; r-&gt;vconst;
		break;

	case OLO:
		v = (uvlong)l-&gt;vconst &lt; (uvlong)r-&gt;vconst;
		break;

	case OLT:
		if(typefd[l-&gt;type-&gt;etype])
			v = l-&gt;fconst &lt; r-&gt;fconst;
		else
			v = l-&gt;vconst &lt; r-&gt;vconst;
		break;

	case OHI:
		v = (uvlong)l-&gt;vconst &gt; (uvlong)r-&gt;vconst;
		break;

	case OGT:
		if(typefd[l-&gt;type-&gt;etype])
			v = l-&gt;fconst &gt; r-&gt;fconst;
		else
			v = l-&gt;vconst &gt; r-&gt;vconst;
		break;

	case OLS:
		v = (uvlong)l-&gt;vconst &lt;= (uvlong)r-&gt;vconst;
		break;

	case OLE:
		if(typefd[l-&gt;type-&gt;etype])
			v = l-&gt;fconst &lt;= r-&gt;fconst;
		else
			v = l-&gt;vconst &lt;= r-&gt;vconst;
		break;

	case OHS:
		v = (uvlong)l-&gt;vconst &gt;= (uvlong)r-&gt;vconst;
		break;

	case OGE:
		if(typefd[l-&gt;type-&gt;etype])
			v = l-&gt;fconst &gt;= r-&gt;fconst;
		else
			v = l-&gt;vconst &gt;= r-&gt;vconst;
		break;

	case OEQ:
		if(typefd[l-&gt;type-&gt;etype])
			v = l-&gt;fconst == r-&gt;fconst;
		else
			v = l-&gt;vconst == r-&gt;vconst;
		break;

	case ONE:
		if(typefd[l-&gt;type-&gt;etype])
			v = l-&gt;fconst != r-&gt;fconst;
		else
			v = l-&gt;vconst != r-&gt;vconst;
		break;

	case ONOT:
		if(typefd[l-&gt;type-&gt;etype])
			v = !l-&gt;fconst;
		else
			v = !l-&gt;vconst;
		break;

	case OANDAND:
		if(typefd[l-&gt;type-&gt;etype])
			v = l-&gt;fconst &amp;&amp; r-&gt;fconst;
		else
			v = l-&gt;vconst &amp;&amp; r-&gt;vconst;
		break;

	case OOROR:
		if(typefd[l-&gt;type-&gt;etype])
			v = l-&gt;fconst || r-&gt;fconst;
		else
			v = l-&gt;vconst || r-&gt;vconst;
		break;
	}
	if(isf) {
		n-&gt;fconst = d;
	} else {
		n-&gt;vconst = convvtox(v, n-&gt;type-&gt;etype);
	}
	n-&gt;oldop = n-&gt;op;
	n-&gt;op = OCONST;
}

void
acom(Node *n)
{
	Type *t;
	Node *l, *r;
	int i;

	switch(n-&gt;op)
	{

	case ONAME:
	case OCONST:
	case OSTRING:
	case OINDREG:
	case OREGISTER:
		return;

	case ONEG:
		l = n-&gt;left;
		if(addo(n) &amp;&amp; addo(l))
			break;
		acom(l);
		return;

	case OADD:
	case OSUB:
	case OMUL:
		l = n-&gt;left;
		r = n-&gt;right;
		if(addo(n)) {
			if(addo(r))
				break;
			if(addo(l))
				break;
		}
		acom(l);
		acom(r);
		return;

	default:
		l = n-&gt;left;
		r = n-&gt;right;
		if(l != Z)
			acom(l);
		if(r != Z)
			acom(r);
		return;
	}

	/* bust terms out */
	t = n-&gt;type;
	term[0].mult = 0;
	term[0].node = Z;
	nterm = 1;
	acom1(1, n);
	if(debug[&#39;m&#39;])
	for(i=0; i&lt;nterm; i++) {
		print(&#34;%d %3lld &#34;, i, term[i].mult);
		prtree1(term[i].node, 1, 0);
	}
	if(nterm &lt; NTERM)
		acom2(n, t);
	n-&gt;type = t;
}

int
acomcmp1(const void *a1, const void *a2)
{
	vlong c1, c2;
	Term *t1, *t2;

	t1 = (Term*)a1;
	t2 = (Term*)a2;
	c1 = t1-&gt;mult;
	if(c1 &lt; 0)
		c1 = -c1;
	c2 = t2-&gt;mult;
	if(c2 &lt; 0)
		c2 = -c2;
	if(c1 &gt; c2)
		return 1;
	if(c1 &lt; c2)
		return -1;
	c1 = 1;
	if(t1-&gt;mult &lt; 0)
		c1 = 0;
	c2 = 1;
	if(t2-&gt;mult &lt; 0)
		c2 = 0;
	if(c2 -= c1)
		return c2;
	if(t2 &gt; t1)
		return 1;
	return -1;
}

int
acomcmp2(const void *a1, const void *a2)
{
	vlong c1, c2;
	Term *t1, *t2;

	t1 = (Term*)a1;
	t2 = (Term*)a2;
	c1 = t1-&gt;mult;
	c2 = t2-&gt;mult;
	if(c1 &gt; c2)
		return 1;
	if(c1 &lt; c2)
		return -1;
	if(t2 &gt; t1)
		return 1;
	return -1;
}

void
acom2(Node *n, Type *t)
{
	Node *l, *r;
	Term trm[NTERM];
	int et, nt, i, j;
	vlong c1, c2;

	/*
	 * copy into automatic
	 */
	c2 = 0;
	nt = nterm;
	for(i=0; i&lt;nt; i++)
		trm[i] = term[i];
	/*
	 * recur on subtrees
	 */
	j = 0;
	for(i=1; i&lt;nt; i++) {
		c1 = trm[i].mult;
		if(c1 == 0)
			continue;
		l = trm[i].node;
		if(l != Z) {
			j = 1;
			acom(l);
		}
	}
	c1 = trm[0].mult;
	if(j == 0) {
		n-&gt;oldop = n-&gt;op;
		n-&gt;op = OCONST;
		n-&gt;vconst = c1;
		return;
	}
	et = t-&gt;etype;

	/*
	 * prepare constant term,
	 * combine it with an addressing term
	 */
	if(c1 != 0) {
		l = new1(OCONST, Z, Z);
		l-&gt;type = t;
		l-&gt;vconst = c1;
		trm[0].mult = 1;
		for(i=1; i&lt;nt; i++) {
			if(trm[i].mult != 1)
				continue;
			r = trm[i].node;
			if(r-&gt;op != OADDR)
				continue;
			r-&gt;type = t;
			l = new1(OADD, r, l);
			l-&gt;type = t;
			trm[i].mult = 0;
			break;
		}
		trm[0].node = l;
	}
	/*
	 * look for factorable terms
	 * c1*i + c1*c2*j -&gt; c1*(i + c2*j)
	 */
	qsort(trm+1, nt-1, sizeof(trm[0]), acomcmp1);
	for(i=nt-1; i&gt;=0; i--) {
		c1 = trm[i].mult;
		if(c1 &lt; 0)
			c1 = -c1;
		if(c1 &lt;= 1)
			continue;
		for(j=i+1; j&lt;nt; j++) {
			c2 = trm[j].mult;
			if(c2 &lt; 0)
				c2 = -c2;
			if(c2 &lt;= 1)
				continue;
			if(c2 % c1)
				continue;
			r = trm[j].node;
			if(r-&gt;type-&gt;etype != et)
				r = acast(t, r);
			c2 = trm[j].mult/trm[i].mult;
			if(c2 != 1 &amp;&amp; c2 != -1) {
				r = new1(OMUL, r, new(OCONST, Z, Z));
				r-&gt;type = t;
				r-&gt;right-&gt;type = t;
				r-&gt;right-&gt;vconst = c2;
			}
			l = trm[i].node;
			if(l-&gt;type-&gt;etype != et)
				l = acast(t, l);
			r = new1(OADD, l, r);
			r-&gt;type = t;
			if(c2 == -1)
				r-&gt;op = OSUB;
			trm[i].node = r;
			trm[j].mult = 0;
		}
	}
	if(debug[&#39;m&#39;]) {
		print(&#34;\n&#34;);
		for(i=0; i&lt;nt; i++) {
			print(&#34;%d %3lld &#34;, i, trm[i].mult);
			prtree1(trm[i].node, 1, 0);
		}
	}

	/*
	 * put it all back together
	 */
	qsort(trm+1, nt-1, sizeof(trm[0]), acomcmp2);
	l = Z;
	for(i=nt-1; i&gt;=0; i--) {
		c1 = trm[i].mult;
		if(c1 == 0)
			continue;
		r = trm[i].node;
		if(r-&gt;type-&gt;etype != et || r-&gt;op == OBIT)
			r = acast(t, r);
		if(c1 != 1 &amp;&amp; c1 != -1) {
			r = new1(OMUL, r, new(OCONST, Z, Z));
			r-&gt;type = t;
			r-&gt;right-&gt;type = t;
			if(c1 &lt; 0) {
				r-&gt;right-&gt;vconst = -c1;
				c1 = -1;
			} else {
				r-&gt;right-&gt;vconst = c1;
				c1 = 1;
			}
		}
		if(l == Z) {
			l = r;
			c2 = c1;
			continue;
		}
		if(c1 &lt; 0)
			if(c2 &lt; 0)
				l = new1(OADD, l, r);
			else
				l = new1(OSUB, l, r);
		else
			if(c2 &lt; 0) {
				l = new1(OSUB, r, l);
				c2 = 1;
			} else
				l = new1(OADD, l, r);
		l-&gt;type = t;
	}
	if(c2 &lt; 0) {
		r = new1(OCONST, 0, 0);
		r-&gt;vconst = 0;
		r-&gt;type = t;
		l = new1(OSUB, r, l);
		l-&gt;type = t;
	}
	*n = *l;
}

void
acom1(vlong v, Node *n)
{
	Node *l, *r;

	if(v == 0 || nterm &gt;= NTERM)
		return;
	if(!addo(n)) {
		if(n-&gt;op == OCONST)
		if(!typefd[n-&gt;type-&gt;etype]) {
			term[0].mult += v*n-&gt;vconst;
			return;
		}
		term[nterm].mult = v;
		term[nterm].node = n;
		nterm++;
		return;
	}
	switch(n-&gt;op) {

	case OCAST:
		acom1(v, n-&gt;left);
		break;

	case ONEG:
		acom1(-v, n-&gt;left);
		break;

	case OADD:
		acom1(v, n-&gt;left);
		acom1(v, n-&gt;right);
		break;

	case OSUB:
		acom1(v, n-&gt;left);
		acom1(-v, n-&gt;right);
		break;

	case OMUL:
		l = n-&gt;left;
		r = n-&gt;right;
		if(l-&gt;op == OCONST)
		if(!typefd[n-&gt;type-&gt;etype]) {
			acom1(v*l-&gt;vconst, r);
			break;
		}
		if(r-&gt;op == OCONST)
		if(!typefd[n-&gt;type-&gt;etype]) {
			acom1(v*r-&gt;vconst, l);
			break;
		}
		break;

	default:
		diag(n, &#34;not addo&#34;);
	}
}

int
addo(Node *n)
{

	if(n != Z)
	if(!typefd[n-&gt;type-&gt;etype])
	if(!typev[n-&gt;type-&gt;etype] || ewidth[TVLONG] == ewidth[TIND])
	switch(n-&gt;op) {

	case OCAST:
		if(nilcast(n-&gt;left-&gt;type, n-&gt;type))
			return 1;
		break;

	case ONEG:
	case OADD:
	case OSUB:
		return 1;

	case OMUL:
		if(n-&gt;left-&gt;op == OCONST)
			return 1;
		if(n-&gt;right-&gt;op == OCONST)
			return 1;
	}
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

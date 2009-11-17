<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5c/mul.c</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5c/mul.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5c/mul.c
// http://code.google.com/p/inferno-os/source/browse/utils/5c/mul.c
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

/*
 * code sequences for multiply by constant.
 * [a-l][0-3]
 *	lsl	$(A-&#39;a&#39;),r0,r1
 * [+][0-7]
 *	add	r0,r1,r2
 * [-][0-7]
 *	sub	r0,r1,r2
 */

static  int	maxmulops = 3;	/* max # of ops to replace mul with */
static	int	multabp;
static	int32	mulval;
static	char*	mulcp;
static	int32	valmax;
static	int	shmax;

static int	docode(char *hp, char *cp, int r0, int r1);
static int	gen1(int len);
static int	gen2(int len, int32 r1);
static int	gen3(int len, int32 r0, int32 r1, int flag);
enum
{
	SR1	= 1&lt;&lt;0,		/* r1 has been shifted */
	SR0	= 1&lt;&lt;1,		/* r0 has been shifted */
	UR1	= 1&lt;&lt;2,		/* r1 has not been used */
	UR0	= 1&lt;&lt;3,		/* r0 has not been used */
};

Multab*
mulcon0(int32 v)
{
	int a1, a2, g;
	Multab *m, *m1;
	char hint[10];

	if(v &lt; 0)
		v = -v;

	/*
	 * look in cache
	 */
	m = multab;
	for(g=0; g&lt;nelem(multab); g++) {
		if(m-&gt;val == v) {
			if(m-&gt;code[0] == 0)
				return 0;
			return m;
		}
		m++;
	}

	/*
	 * select a spot in cache to overwrite
	 */
	multabp++;
	if(multabp &lt; 0 || multabp &gt;= nelem(multab))
		multabp = 0;
	m = multab+multabp;
	m-&gt;val = v;
	mulval = v;

	/*
	 * look in execption hint table
	 */
	a1 = 0;
	a2 = hintabsize;
	for(;;) {
		if(a1 &gt;= a2)
			goto no;
		g = (a2 + a1)/2;
		if(v &lt; hintab[g].val) {
			a2 = g;
			continue;
		}
		if(v &gt; hintab[g].val) {
			a1 = g+1;
			continue;
		}
		break;
	}

	if(docode(hintab[g].hint, m-&gt;code, 1, 0))
		return m;
	print(&#34;multiply table failure %ld\n&#34;, v);
	m-&gt;code[0] = 0;
	return 0;

no:
	/*
	 * try to search
	 */
	hint[0] = 0;
	for(g=1; g&lt;=maxmulops; g++) {
		if(g &gt;= maxmulops &amp;&amp; v &gt;= 65535)
			break;
		mulcp = hint+g;
		*mulcp = 0;
		if(gen1(g)) {
			if(docode(hint, m-&gt;code, 1, 0))
				return m;
			print(&#34;multiply table failure %ld\n&#34;, v);
			break;
		}
	}

	/*
	 * try a recur followed by a shift
	 */
	g = 0;
	while(!(v &amp; 1)) {
		g++;
		v &gt;&gt;= 1;
	}
	if(g) {
		m1 = mulcon0(v);
		if(m1) {
			strcpy(m-&gt;code, m1-&gt;code);
			sprint(strchr(m-&gt;code, 0), &#34;%c0&#34;, g+&#39;a&#39;);
			return m;
		}
	}
	m-&gt;code[0] = 0;
	return 0;
}

static int
docode(char *hp, char *cp, int r0, int r1)
{
	int c, i;

	c = *hp++;
	*cp = c;
	cp += 2;
	switch(c) {
	default:
		c -= &#39;a&#39;;
		if(c &lt; 1 || c &gt;= 30)
			break;
		for(i=0; i&lt;4; i++) {
			switch(i) {
			case 0:
				if(docode(hp, cp, r0&lt;&lt;c, r1))
					goto out;
				break;
			case 1:
				if(docode(hp, cp, r1&lt;&lt;c, r1))
					goto out;
				break;
			case 2:
				if(docode(hp, cp, r0, r0&lt;&lt;c))
					goto out;
				break;
			case 3:
				if(docode(hp, cp, r0, r1&lt;&lt;c))
					goto out;
				break;
			}
		}
		break;

	case &#39;+&#39;:
		for(i=0; i&lt;8; i++) {
			cp[-1] = i+&#39;0&#39;;
			switch(i) {
			case 1:
				if(docode(hp, cp, r0+r1, r1))
					goto out;
				break;
			case 5:
				if(docode(hp, cp, r0, r0+r1))
					goto out;
				break;
			}
		}
		break;

	case &#39;-&#39;:
		for(i=0; i&lt;8; i++) {
			cp[-1] = i+&#39;0&#39;;
			switch(i) {
			case 1:
				if(docode(hp, cp, r0-r1, r1))
					goto out;
				break;
			case 2:
				if(docode(hp, cp, r1-r0, r1))
					goto out;
				break;
			case 5:
				if(docode(hp, cp, r0, r0-r1))
					goto out;
				break;
			case 6:
				if(docode(hp, cp, r0, r1-r0))
					goto out;
				break;
			}
		}
		break;

	case 0:
		if(r0 == mulval)
			return 1;
	}
	return 0;

out:
	cp[-1] = i+&#39;0&#39;;
	return 1;
}

static int
gen1(int len)
{
	int i;

	for(shmax=1; shmax&lt;30; shmax++) {
		valmax = 1&lt;&lt;shmax;
		if(valmax &gt;= mulval)
			break;
	}
	if(mulval == 1)
		return 1;

	len--;
	for(i=1; i&lt;=shmax; i++)
		if(gen2(len, 1&lt;&lt;i)) {
			*--mulcp = &#39;a&#39;+i;
			return 1;
		}
	return 0;
}

static int
gen2(int len, int32 r1)
{
	int i;

	if(len &lt;= 0) {
		if(r1 == mulval)
			return 1;
		return 0;
	}

	len--;
	if(len == 0)
		goto calcr0;

	if(gen3(len, r1, r1+1, UR1)) {
		i = &#39;+&#39;;
		goto out;
	}
	if(gen3(len, r1-1, r1, UR0)) {
		i = &#39;-&#39;;
		goto out;
	}
	if(gen3(len, 1, r1+1, UR1)) {
		i = &#39;+&#39;;
		goto out;
	}
	if(gen3(len, 1, r1-1, UR1)) {
		i = &#39;-&#39;;
		goto out;
	}

	return 0;

calcr0:
	if(mulval == r1+1) {
		i = &#39;+&#39;;
		goto out;
	}
	if(mulval == r1-1) {
		i = &#39;-&#39;;
		goto out;
	}
	return 0;

out:
	*--mulcp = i;
	return 1;
}

static int
gen3(int len, int32 r0, int32 r1, int flag)
{
	int i, f1, f2;
	int32 x;

	if(r0 &lt;= 0 ||
	   r0 &gt;= r1 ||
	   r1 &gt; valmax)
		return 0;

	len--;
	if(len == 0)
		goto calcr0;

	if(!(flag &amp; UR1)) {
		f1 = UR1|SR1;
		for(i=1; i&lt;=shmax; i++) {
			x = r0&lt;&lt;i;
			if(x &gt; valmax)
				break;
			if(gen3(len, r0, x, f1)) {
				i += &#39;a&#39;;
				goto out;
			}
		}
	}

	if(!(flag &amp; UR0)) {
		f1 = UR1|SR1;
		for(i=1; i&lt;=shmax; i++) {
			x = r1&lt;&lt;i;
			if(x &gt; valmax)
				break;
			if(gen3(len, r1, x, f1)) {
				i += &#39;a&#39;;
				goto out;
			}
		}
	}

	if(!(flag &amp; SR1)) {
		f1 = UR1|SR1|(flag&amp;UR0);
		for(i=1; i&lt;=shmax; i++) {
			x = r1&lt;&lt;i;
			if(x &gt; valmax)
				break;
			if(gen3(len, r0, x, f1)) {
				i += &#39;a&#39;;
				goto out;
			}
		}
	}

	if(!(flag &amp; SR0)) {
		f1 = UR0|SR0|(flag&amp;(SR1|UR1));

		f2 = UR1|SR1;
		if(flag &amp; UR1)
			f2 |= UR0;
		if(flag &amp; SR1)
			f2 |= SR0;

		for(i=1; i&lt;=shmax; i++) {
			x = r0&lt;&lt;i;
			if(x &gt; valmax)
				break;
			if(x &gt; r1) {
				if(gen3(len, r1, x, f2)) {
					i += &#39;a&#39;;
					goto out;
				}
			} else
				if(gen3(len, x, r1, f1)) {
					i += &#39;a&#39;;
					goto out;
				}
		}
	}

	x = r1+r0;
	if(gen3(len, r0, x, UR1)) {
		i = &#39;+&#39;;
		goto out;
	}

	if(gen3(len, r1, x, UR1)) {
		i = &#39;+&#39;;
		goto out;
	}

	x = r1-r0;
	if(gen3(len, x, r1, UR0)) {
		i = &#39;-&#39;;
		goto out;
	}

	if(x &gt; r0) {
		if(gen3(len, r0, x, UR1)) {
			i = &#39;-&#39;;
			goto out;
		}
	} else
		if(gen3(len, x, r0, UR0)) {
			i = &#39;-&#39;;
			goto out;
		}

	return 0;

calcr0:
	f1 = flag &amp; (UR0|UR1);
	if(f1 == UR1) {
		for(i=1; i&lt;=shmax; i++) {
			x = r1&lt;&lt;i;
			if(x &gt;= mulval) {
				if(x == mulval) {
					i += &#39;a&#39;;
					goto out;
				}
				break;
			}
		}
	}

	if(mulval == r1+r0) {
		i = &#39;+&#39;;
		goto out;
	}
	if(mulval == r1-r0) {
		i = &#39;-&#39;;
		goto out;
	}

	return 0;

out:
	*--mulcp = i;
	return 1;
}

/*
 * hint table has numbers that
 * the search algorithm fails on.
 * &lt;1000:
 *	all numbers
 * &lt;5000:
 * 	÷ by 5
 * &lt;10000:
 * 	÷ by 50
 * &lt;65536:
 * 	÷ by 250
 */
Hintab	hintab[] =
{
	683,	&#34;b++d+e+&#34;,
	687,	&#34;b+e++e-&#34;,
	691,	&#34;b++d+e+&#34;,
	731,	&#34;b++d+e+&#34;,
	811,	&#34;b++d+i+&#34;,
	821,	&#34;b++e+e+&#34;,
	843,	&#34;b+d++e+&#34;,
	851,	&#34;b+f-+e-&#34;,
	853,	&#34;b++e+e+&#34;,
	877,	&#34;c++++g-&#34;,
	933,	&#34;b+c++g-&#34;,
	981,	&#34;c-+e-d+&#34;,
	1375,	&#34;b+c+b+h-&#34;,
	1675,	&#34;d+b++h+&#34;,
	2425,	&#34;c++f-e+&#34;,
	2675,	&#34;c+d++f-&#34;,
	2750,	&#34;b+d-b+h-&#34;,
	2775,	&#34;c-+g-e-&#34;,
	3125,	&#34;b++e+g+&#34;,
	3275,	&#34;b+c+g+e+&#34;,
	3350,	&#34;c++++i+&#34;,
	3475,	&#34;c-+e-f-&#34;,
	3525,	&#34;c-+d+g-&#34;,
	3625,	&#34;c-+e-j+&#34;,
	3675,	&#34;b+d+d+e+&#34;,
	3725,	&#34;b+d-+h+&#34;,
	3925,	&#34;b+d+f-d-&#34;,
	4275,	&#34;b+g++e+&#34;,
	4325,	&#34;b+h-+d+&#34;,
	4425,	&#34;b+b+g-j-&#34;,
	4525,	&#34;b+d-d+f+&#34;,
	4675,	&#34;c++d-g+&#34;,
	4775,	&#34;b+d+b+g-&#34;,
	4825,	&#34;c+c-+i-&#34;,
	4850,	&#34;c++++i-&#34;,
	4925,	&#34;b++e-g-&#34;,
	4975,	&#34;c+f++e-&#34;,
	5500,	&#34;b+g-c+d+&#34;,
	6700,	&#34;d+b++i+&#34;,
	9700,	&#34;d++++j-&#34;,
	11000,	&#34;b+f-c-h-&#34;,
	11750,	&#34;b+d+g+j-&#34;,
	12500,	&#34;b+c+e-k+&#34;,
	13250,	&#34;b+d+e-f+&#34;,
	13750,	&#34;b+h-c-d+&#34;,
	14250,	&#34;b+g-c+e-&#34;,
	14500,	&#34;c+f+j-d-&#34;,
	14750,	&#34;d-g--f+&#34;,
	16750,	&#34;b+e-d-n+&#34;,
	17750,	&#34;c+h-b+e+&#34;,
	18250,	&#34;d+b+h-d+&#34;,
	18750,	&#34;b+g-++f+&#34;,
	19250,	&#34;b+e+b+h+&#34;,
	19750,	&#34;b++h--f-&#34;,
	20250,	&#34;b+e-l-c+&#34;,
	20750,	&#34;c++bi+e-&#34;,
	21250,	&#34;b+i+l+c+&#34;,
	22000,	&#34;b+e+d-g-&#34;,
	22250,	&#34;b+d-h+k-&#34;,
	22750,	&#34;b+d-e-g+&#34;,
	23250,	&#34;b+c+h+e-&#34;,
	23500,	&#34;b+g-c-g-&#34;,
	23750,	&#34;b+g-b+h-&#34;,
	24250,	&#34;c++g+m-&#34;,
	24750,	&#34;b+e+e+j-&#34;,
	25000,	&#34;b++dh+g+&#34;,
	25250,	&#34;b+e+d-g-&#34;,
	25750,	&#34;b+e+b+j+&#34;,
	26250,	&#34;b+h+c+e+&#34;,
	26500,	&#34;b+h+c+g+&#34;,
	26750,	&#34;b+d+e+g-&#34;,
	27250,	&#34;b+e+e+f+&#34;,
	27500,	&#34;c-i-c-d+&#34;,
	27750,	&#34;b+bd++j+&#34;,
	28250,	&#34;d-d-++i-&#34;,
	28500,	&#34;c+c-h-e-&#34;,
	29000,	&#34;b+g-d-f+&#34;,
	29500,	&#34;c+h+++e-&#34;,
	29750,	&#34;b+g+f-c+&#34;,
	30250,	&#34;b+f-g-c+&#34;,
	33500,	&#34;c-f-d-n+&#34;,
	33750,	&#34;b+d-b+j-&#34;,
	34250,	&#34;c+e+++i+&#34;,
	35250,	&#34;e+b+d+k+&#34;,
	35500,	&#34;c+e+d-g-&#34;,
	35750,	&#34;c+i-++e+&#34;,
	36250,	&#34;b+bh-d+e+&#34;,
	36500,	&#34;c+c-h-e-&#34;,
	36750,	&#34;d+e--i+&#34;,
	37250,	&#34;b+g+g+b+&#34;,
	37500,	&#34;b+h-b+f+&#34;,
	37750,	&#34;c+be++j-&#34;,
	38500,	&#34;b+e+b+i+&#34;,
	38750,	&#34;d+i-b+d+&#34;,
	39250,	&#34;b+g-l-+d+&#34;,
	39500,	&#34;b+g-c+g-&#34;,
	39750,	&#34;b+bh-c+f-&#34;,
	40250,	&#34;b+bf+d+g-&#34;,
	40500,	&#34;b+g-c+g+&#34;,
	40750,	&#34;c+b+i-e+&#34;,
	41250,	&#34;d++bf+h+&#34;,
	41500,	&#34;b+j+c+d-&#34;,
	41750,	&#34;c+f+b+h-&#34;,
	42500,	&#34;c+h++g+&#34;,
	42750,	&#34;b+g+d-f-&#34;,
	43250,	&#34;b+l-e+d-&#34;,
	43750,	&#34;c+bd+h+f-&#34;,
	44000,	&#34;b+f+g-d-&#34;,
	44250,	&#34;b+d-g--f+&#34;,
	44500,	&#34;c+e+c+h+&#34;,
	44750,	&#34;b+e+d-h-&#34;,
	45250,	&#34;b++g+j-g+&#34;,
	45500,	&#34;c+d+e-g+&#34;,
	45750,	&#34;b+d-h-e-&#34;,
	46250,	&#34;c+bd++j+&#34;,
	46500,	&#34;b+d-c-j-&#34;,
	46750,	&#34;e-e-b+g-&#34;,
	47000,	&#34;b+c+d-j-&#34;,
	47250,	&#34;b+e+e-g-&#34;,
	47500,	&#34;b+g-c-h-&#34;,
	47750,	&#34;b+f-c+h-&#34;,
	48250,	&#34;d--h+n-&#34;,
	48500,	&#34;b+c-g+m-&#34;,
	48750,	&#34;b+e+e-g+&#34;,
	49500,	&#34;c-f+e+j-&#34;,
	49750,	&#34;c+c+g++f-&#34;,
	50000,	&#34;b+e+e+k+&#34;,
	50250,	&#34;b++i++g+&#34;,
	50500,	&#34;c+g+f-i+&#34;,
	50750,	&#34;b+e+d+k-&#34;,
	51500,	&#34;b+i+c-f+&#34;,
	51750,	&#34;b+bd+g-e-&#34;,
	52250,	&#34;b+d+g-j+&#34;,
	52500,	&#34;c+c+f+g+&#34;,
	52750,	&#34;b+c+e+i+&#34;,
	53000,	&#34;b+i+c+g+&#34;,
	53500,	&#34;c+g+g-n+&#34;,
	53750,	&#34;b+j+d-c+&#34;,
	54250,	&#34;b+d-g-j-&#34;,
	54500,	&#34;c-f+e+f+&#34;,
	54750,	&#34;b+f-+c+g+&#34;,
	55000,	&#34;b+g-d-g-&#34;,
	55250,	&#34;b+e+e+g+&#34;,
	55500,	&#34;b+cd++j+&#34;,
	55750,	&#34;b+bh-d-f-&#34;,
	56250,	&#34;c+d-b+j-&#34;,
	56500,	&#34;c+d+c+i+&#34;,
	56750,	&#34;b+e+d++h-&#34;,
	57000,	&#34;b+d+g-f+&#34;,
	57250,	&#34;b+f-m+d-&#34;,
	57750,	&#34;b+i+c+e-&#34;,
	58000,	&#34;b+e+d+h+&#34;,
	58250,	&#34;c+b+g+g+&#34;,
	58750,	&#34;d-e-j--e+&#34;,
	59000,	&#34;d-i-+e+&#34;,
	59250,	&#34;e--h-m+&#34;,
	59500,	&#34;c+c-h+f-&#34;,
	59750,	&#34;b+bh-e+i-&#34;,
	60250,	&#34;b+bh-e-e-&#34;,
	60500,	&#34;c+c-g-g-&#34;,
	60750,	&#34;b+e-l-e-&#34;,
	61250,	&#34;b+g-g-c+&#34;,
	61750,	&#34;b+g-c+g+&#34;,
	62250,	&#34;f--+c-i-&#34;,
	62750,	&#34;e+f--+g+&#34;,
	64750,	&#34;b+f+d+p-&#34;,
};
int	hintabsize	= nelem(hintab);
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

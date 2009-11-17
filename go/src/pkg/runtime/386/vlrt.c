<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/386/vlrt.c</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/386/vlrt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno&#39;s libkern/vlrt-386.c
// http://code.google.com/p/inferno-os/source/browse/libkern/vlrt-386.c
//
//         Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//         Revisions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).  All rights reserved.
//         Portions Copyright 2009 The Go Authors. All rights reserved.
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

/*
 * C runtime for 64-bit divide, others.
 *
 * TODO(rsc): The simple functions are dregs--8c knows how
 * to generate the code directly now.  Find and remove.
 */

typedef	unsigned long	ulong;
typedef	unsigned int	uint;
typedef	unsigned short	ushort;
typedef	unsigned char	uchar;
typedef	signed char	schar;

#define	SIGN(n)	(1UL&lt;&lt;(n-1))

typedef	struct	Vlong	Vlong;
struct	Vlong
{
	union
	{
		long long	v;
		struct
		{
			ulong	lo;
			ulong	hi;
		};
		struct
		{
			ushort	lols;
			ushort	loms;
			ushort	hils;
			ushort	hims;
		};
	};
};

void	abort(void);

void
_d2v(Vlong *y, double d)
{
	union { double d; struct Vlong; } x;
	ulong xhi, xlo, ylo, yhi;
	int sh;

	x.d = d;

	xhi = (x.hi &amp; 0xfffff) | 0x100000;
	xlo = x.lo;
	sh = 1075 - ((x.hi &gt;&gt; 20) &amp; 0x7ff);

	ylo = 0;
	yhi = 0;
	if(sh &gt;= 0) {
		/* v = (hi||lo) &gt;&gt; sh */
		if(sh &lt; 32) {
			if(sh == 0) {
				ylo = xlo;
				yhi = xhi;
			} else {
				ylo = (xlo &gt;&gt; sh) | (xhi &lt;&lt; (32-sh));
				yhi = xhi &gt;&gt; sh;
			}
		} else {
			if(sh == 32) {
				ylo = xhi;
			} else
			if(sh &lt; 64) {
				ylo = xhi &gt;&gt; (sh-32);
			}
		}
	} else {
		/* v = (hi||lo) &lt;&lt; -sh */
		sh = -sh;
		if(sh &lt;= 10) {
			ylo = xlo &lt;&lt; sh;
			yhi = (xhi &lt;&lt; sh) | (xlo &gt;&gt; (32-sh));
		} else {
			/* overflow */
			yhi = d;	/* causes something awful */
		}
	}
	if(x.hi &amp; SIGN(32)) {
		if(ylo != 0) {
			ylo = -ylo;
			yhi = ~yhi;
		} else
			yhi = -yhi;
	}

	y-&gt;hi = yhi;
	y-&gt;lo = ylo;
}

void
_f2v(Vlong *y, float f)
{

	_d2v(y, f);
}

double
_v2d(Vlong x)
{
	if(x.hi &amp; SIGN(32)) {
		if(x.lo) {
			x.lo = -x.lo;
			x.hi = ~x.hi;
		} else
			x.hi = -x.hi;
		return -((long)x.hi*4294967296. + x.lo);
	}
	return (long)x.hi*4294967296. + x.lo;
}

float
_v2f(Vlong x)
{
	return _v2d(x);
}

ulong	_div64by32(Vlong, ulong, ulong*);
void	_mul64by32(Vlong*, Vlong, ulong);

static void
slowdodiv(Vlong num, Vlong den, Vlong *q, Vlong *r)
{
	ulong numlo, numhi, denhi, denlo, quohi, quolo, t;
	int i;

	numhi = num.hi;
	numlo = num.lo;
	denhi = den.hi;
	denlo = den.lo;

	/*
	 * get a divide by zero
	 */
	if(denlo==0 &amp;&amp; denhi==0) {
		numlo = numlo / denlo;
	}

	/*
	 * set up the divisor and find the number of iterations needed
	 */
	if(numhi &gt;= SIGN(32)) {
		quohi = SIGN(32);
		quolo = 0;
	} else {
		quohi = numhi;
		quolo = numlo;
	}
	i = 0;
	while(denhi &lt; quohi || (denhi == quohi &amp;&amp; denlo &lt; quolo)) {
		denhi = (denhi&lt;&lt;1) | (denlo&gt;&gt;31);
		denlo &lt;&lt;= 1;
		i++;
	}

	quohi = 0;
	quolo = 0;
	for(; i &gt;= 0; i--) {
		quohi = (quohi&lt;&lt;1) | (quolo&gt;&gt;31);
		quolo &lt;&lt;= 1;
		if(numhi &gt; denhi || (numhi == denhi &amp;&amp; numlo &gt;= denlo)) {
			t = numlo;
			numlo -= denlo;
			if(numlo &gt; t)
				numhi--;
			numhi -= denhi;
			quolo |= 1;
		}
		denlo = (denlo&gt;&gt;1) | (denhi&lt;&lt;31);
		denhi &gt;&gt;= 1;
	}

	if(q) {
		q-&gt;lo = quolo;
		q-&gt;hi = quohi;
	}
	if(r) {
		r-&gt;lo = numlo;
		r-&gt;hi = numhi;
	}
}

static void
dodiv(Vlong num, Vlong den, Vlong *qp, Vlong *rp)
{
	ulong n;
	Vlong x, q, r;

	if(den.hi &gt; num.hi || (den.hi == num.hi &amp;&amp; den.lo &gt; num.lo)){
		if(qp) {
			qp-&gt;hi = 0;
			qp-&gt;lo = 0;
		}
		if(rp) {
			rp-&gt;hi = num.hi;
			rp-&gt;lo = num.lo;
		}
		return;
	}

	if(den.hi != 0){
		q.hi = 0;
		n = num.hi/den.hi;
		_mul64by32(&amp;x, den, n);
		if(x.hi &gt; num.hi || (x.hi == num.hi &amp;&amp; x.lo &gt; num.lo))
			slowdodiv(num, den, &amp;q, &amp;r);
		else {
			q.lo = n;
			r.v = num.v - x.v;
		}
	} else {
		if(num.hi &gt;= den.lo){
			q.hi = n = num.hi/den.lo;
			num.hi -= den.lo*n;
		} else {
			q.hi = 0;
		}
		q.lo = _div64by32(num, den.lo, &amp;r.lo);
		r.hi = 0;
	}
	if(qp) {
		qp-&gt;lo = q.lo;
		qp-&gt;hi = q.hi;
	}
	if(rp) {
		rp-&gt;lo = r.lo;
		rp-&gt;hi = r.hi;
	}
}

void
_divvu(Vlong *q, Vlong n, Vlong d)
{

	if(n.hi == 0 &amp;&amp; d.hi == 0) {
		q-&gt;hi = 0;
		q-&gt;lo = n.lo / d.lo;
		return;
	}
	dodiv(n, d, q, 0);
}

void
runtime·uint64div(Vlong n, Vlong d, Vlong q)
{
	_divvu(&amp;q, n, d);
}

void
_modvu(Vlong *r, Vlong n, Vlong d)
{

	if(n.hi == 0 &amp;&amp; d.hi == 0) {
		r-&gt;hi = 0;
		r-&gt;lo = n.lo % d.lo;
		return;
	}
	dodiv(n, d, 0, r);
}

void
runtime·uint64mod(Vlong n, Vlong d, Vlong q)
{
	_modvu(&amp;q, n, d);
}

static void
vneg(Vlong *v)
{

	if(v-&gt;lo == 0) {
		v-&gt;hi = -v-&gt;hi;
		return;
	}
	v-&gt;lo = -v-&gt;lo;
	v-&gt;hi = ~v-&gt;hi;
}

void
_divv(Vlong *q, Vlong n, Vlong d)
{
	long nneg, dneg;

	if(n.hi == (((long)n.lo)&gt;&gt;31) &amp;&amp; d.hi == (((long)d.lo)&gt;&gt;31)) {
		if((long)n.lo == -0x80000000 &amp;&amp; (long)d.lo == -1) {
			// special case: 32-bit -0x80000000 / -1 causes divide error,
			// but it&#39;s okay in this 64-bit context.
			q-&gt;lo = 0x80000000;
			q-&gt;hi = 0;
			return;
		}
		q-&gt;lo = (long)n.lo / (long)d.lo;
		q-&gt;hi = ((long)q-&gt;lo) &gt;&gt; 31;
		return;
	}
	nneg = n.hi &gt;&gt; 31;
	if(nneg)
		vneg(&amp;n);
	dneg = d.hi &gt;&gt; 31;
	if(dneg)
		vneg(&amp;d);
	dodiv(n, d, q, 0);
	if(nneg != dneg)
		vneg(q);
}

void
runtime·int64div(Vlong n, Vlong d, Vlong q)
{
	_divv(&amp;q, n, d);
}

void
_modv(Vlong *r, Vlong n, Vlong d)
{
	long nneg, dneg;

	if(n.hi == (((long)n.lo)&gt;&gt;31) &amp;&amp; d.hi == (((long)d.lo)&gt;&gt;31)) {
		if((long)n.lo == -0x80000000 &amp;&amp; (long)d.lo == -1) {
			// special case: 32-bit -0x80000000 % -1 causes divide error,
			// but it&#39;s okay in this 64-bit context.
			r-&gt;lo = 0;
			r-&gt;hi = 0;
			return;
		}
		r-&gt;lo = (long)n.lo % (long)d.lo;
		r-&gt;hi = ((long)r-&gt;lo) &gt;&gt; 31;
		return;
	}
	nneg = n.hi &gt;&gt; 31;
	if(nneg)
		vneg(&amp;n);
	dneg = d.hi &gt;&gt; 31;
	if(dneg)
		vneg(&amp;d);
	dodiv(n, d, 0, r);
	if(nneg)
		vneg(r);
}

void
runtime·int64mod(Vlong n, Vlong d, Vlong q)
{
	_modv(&amp;q, n, d);
}

void
_rshav(Vlong *r, Vlong a, int b)
{
	long t;

	t = a.hi;
	if(b &gt;= 32) {
		r-&gt;hi = t&gt;&gt;31;
		if(b &gt;= 64) {
			/* this is illegal re C standard */
			r-&gt;lo = t&gt;&gt;31;
			return;
		}
		r-&gt;lo = t &gt;&gt; (b-32);
		return;
	}
	if(b &lt;= 0) {
		r-&gt;hi = t;
		r-&gt;lo = a.lo;
		return;
	}
	r-&gt;hi = t &gt;&gt; b;
	r-&gt;lo = (t &lt;&lt; (32-b)) | (a.lo &gt;&gt; b);
}

void
_rshlv(Vlong *r, Vlong a, int b)
{
	ulong t;

	t = a.hi;
	if(b &gt;= 32) {
		r-&gt;hi = 0;
		if(b &gt;= 64) {
			/* this is illegal re C standard */
			r-&gt;lo = 0;
			return;
		}
		r-&gt;lo = t &gt;&gt; (b-32);
		return;
	}
	if(b &lt;= 0) {
		r-&gt;hi = t;
		r-&gt;lo = a.lo;
		return;
	}
	r-&gt;hi = t &gt;&gt; b;
	r-&gt;lo = (t &lt;&lt; (32-b)) | (a.lo &gt;&gt; b);
}

void
_lshv(Vlong *r, Vlong a, int b)
{
	ulong t;

	t = a.lo;
	if(b &gt;= 32) {
		r-&gt;lo = 0;
		if(b &gt;= 64) {
			/* this is illegal re C standard */
			r-&gt;hi = 0;
			return;
		}
		r-&gt;hi = t &lt;&lt; (b-32);
		return;
	}
	if(b &lt;= 0) {
		r-&gt;lo = t;
		r-&gt;hi = a.hi;
		return;
	}
	r-&gt;lo = t &lt;&lt; b;
	r-&gt;hi = (t &gt;&gt; (32-b)) | (a.hi &lt;&lt; b);
}

void
_andv(Vlong *r, Vlong a, Vlong b)
{
	r-&gt;hi = a.hi &amp; b.hi;
	r-&gt;lo = a.lo &amp; b.lo;
}

void
_orv(Vlong *r, Vlong a, Vlong b)
{
	r-&gt;hi = a.hi | b.hi;
	r-&gt;lo = a.lo | b.lo;
}

void
_xorv(Vlong *r, Vlong a, Vlong b)
{
	r-&gt;hi = a.hi ^ b.hi;
	r-&gt;lo = a.lo ^ b.lo;
}

void
_vpp(Vlong *l, Vlong *r)
{

	l-&gt;hi = r-&gt;hi;
	l-&gt;lo = r-&gt;lo;
	r-&gt;lo++;
	if(r-&gt;lo == 0)
		r-&gt;hi++;
}

void
_vmm(Vlong *l, Vlong *r)
{

	l-&gt;hi = r-&gt;hi;
	l-&gt;lo = r-&gt;lo;
	if(r-&gt;lo == 0)
		r-&gt;hi--;
	r-&gt;lo--;
}

void
_ppv(Vlong *l, Vlong *r)
{

	r-&gt;lo++;
	if(r-&gt;lo == 0)
		r-&gt;hi++;
	l-&gt;hi = r-&gt;hi;
	l-&gt;lo = r-&gt;lo;
}

void
_mmv(Vlong *l, Vlong *r)
{

	if(r-&gt;lo == 0)
		r-&gt;hi--;
	r-&gt;lo--;
	l-&gt;hi = r-&gt;hi;
	l-&gt;lo = r-&gt;lo;
}

void
_vasop(Vlong *ret, void *lv, void fn(Vlong*, Vlong, Vlong), int type, Vlong rv)
{
	Vlong t, u;

	u.lo = 0;
	u.hi = 0;
	switch(type) {
	default:
		abort();
		break;

	case 1:	/* schar */
		t.lo = *(schar*)lv;
		t.hi = t.lo &gt;&gt; 31;
		fn(&amp;u, t, rv);
		*(schar*)lv = u.lo;
		break;

	case 2:	/* uchar */
		t.lo = *(uchar*)lv;
		t.hi = 0;
		fn(&amp;u, t, rv);
		*(uchar*)lv = u.lo;
		break;

	case 3:	/* short */
		t.lo = *(short*)lv;
		t.hi = t.lo &gt;&gt; 31;
		fn(&amp;u, t, rv);
		*(short*)lv = u.lo;
		break;

	case 4:	/* ushort */
		t.lo = *(ushort*)lv;
		t.hi = 0;
		fn(&amp;u, t, rv);
		*(ushort*)lv = u.lo;
		break;

	case 9:	/* int */
		t.lo = *(int*)lv;
		t.hi = t.lo &gt;&gt; 31;
		fn(&amp;u, t, rv);
		*(int*)lv = u.lo;
		break;

	case 10:	/* uint */
		t.lo = *(uint*)lv;
		t.hi = 0;
		fn(&amp;u, t, rv);
		*(uint*)lv = u.lo;
		break;

	case 5:	/* long */
		t.lo = *(long*)lv;
		t.hi = t.lo &gt;&gt; 31;
		fn(&amp;u, t, rv);
		*(long*)lv = u.lo;
		break;

	case 6:	/* ulong */
		t.lo = *(ulong*)lv;
		t.hi = 0;
		fn(&amp;u, t, rv);
		*(ulong*)lv = u.lo;
		break;

	case 7:	/* vlong */
	case 8:	/* uvlong */
		fn(&amp;u, *(Vlong*)lv, rv);
		*(Vlong*)lv = u;
		break;
	}
	*ret = u;
}

void
_p2v(Vlong *ret, void *p)
{
	long t;

	t = (ulong)p;
	ret-&gt;lo = t;
	ret-&gt;hi = 0;
}

void
_sl2v(Vlong *ret, long sl)
{
	long t;

	t = sl;
	ret-&gt;lo = t;
	ret-&gt;hi = t &gt;&gt; 31;
}

void
_ul2v(Vlong *ret, ulong ul)
{
	long t;

	t = ul;
	ret-&gt;lo = t;
	ret-&gt;hi = 0;
}

void
_si2v(Vlong *ret, int si)
{
	long t;

	t = si;
	ret-&gt;lo = t;
	ret-&gt;hi = t &gt;&gt; 31;
}

void
_ui2v(Vlong *ret, uint ui)
{
	long t;

	t = ui;
	ret-&gt;lo = t;
	ret-&gt;hi = 0;
}

void
_sh2v(Vlong *ret, long sh)
{
	long t;

	t = (sh &lt;&lt; 16) &gt;&gt; 16;
	ret-&gt;lo = t;
	ret-&gt;hi = t &gt;&gt; 31;
}

void
_uh2v(Vlong *ret, ulong ul)
{
	long t;

	t = ul &amp; 0xffff;
	ret-&gt;lo = t;
	ret-&gt;hi = 0;
}

void
_sc2v(Vlong *ret, long uc)
{
	long t;

	t = (uc &lt;&lt; 24) &gt;&gt; 24;
	ret-&gt;lo = t;
	ret-&gt;hi = t &gt;&gt; 31;
}

void
_uc2v(Vlong *ret, ulong ul)
{
	long t;

	t = ul &amp; 0xff;
	ret-&gt;lo = t;
	ret-&gt;hi = 0;
}

long
_v2sc(Vlong rv)
{
	long t;

	t = rv.lo &amp; 0xff;
	return (t &lt;&lt; 24) &gt;&gt; 24;
}

long
_v2uc(Vlong rv)
{

	return rv.lo &amp; 0xff;
}

long
_v2sh(Vlong rv)
{
	long t;

	t = rv.lo &amp; 0xffff;
	return (t &lt;&lt; 16) &gt;&gt; 16;
}

long
_v2uh(Vlong rv)
{

	return rv.lo &amp; 0xffff;
}

long
_v2sl(Vlong rv)
{

	return rv.lo;
}

long
_v2ul(Vlong rv)
{

	return rv.lo;
}

long
_v2si(Vlong rv)
{

	return rv.lo;
}

long
_v2ui(Vlong rv)
{

	return rv.lo;
}

int
_testv(Vlong rv)
{
	return rv.lo || rv.hi;
}

int
_eqv(Vlong lv, Vlong rv)
{
	return lv.lo == rv.lo &amp;&amp; lv.hi == rv.hi;
}

int
_nev(Vlong lv, Vlong rv)
{
	return lv.lo != rv.lo || lv.hi != rv.hi;
}

int
_ltv(Vlong lv, Vlong rv)
{
	return (long)lv.hi &lt; (long)rv.hi ||
		(lv.hi == rv.hi &amp;&amp; lv.lo &lt; rv.lo);
}

int
_lev(Vlong lv, Vlong rv)
{
	return (long)lv.hi &lt; (long)rv.hi ||
		(lv.hi == rv.hi &amp;&amp; lv.lo &lt;= rv.lo);
}

int
_gtv(Vlong lv, Vlong rv)
{
	return (long)lv.hi &gt; (long)rv.hi ||
		(lv.hi == rv.hi &amp;&amp; lv.lo &gt; rv.lo);
}

int
_gev(Vlong lv, Vlong rv)
{
	return (long)lv.hi &gt; (long)rv.hi ||
		(lv.hi == rv.hi &amp;&amp; lv.lo &gt;= rv.lo);
}

int
_lov(Vlong lv, Vlong rv)
{
	return lv.hi &lt; rv.hi ||
		(lv.hi == rv.hi &amp;&amp; lv.lo &lt; rv.lo);
}

int
_lsv(Vlong lv, Vlong rv)
{
	return lv.hi &lt; rv.hi ||
		(lv.hi == rv.hi &amp;&amp; lv.lo &lt;= rv.lo);
}

int
_hiv(Vlong lv, Vlong rv)
{
	return lv.hi &gt; rv.hi ||
		(lv.hi == rv.hi &amp;&amp; lv.lo &gt; rv.lo);
}

int
_hsv(Vlong lv, Vlong rv)
{
	return lv.hi &gt; rv.hi ||
		(lv.hi == rv.hi &amp;&amp; lv.lo &gt;= rv.lo);
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5c/sgen.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/5c/sgen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5c/sgen.c
// http://code.google.com/p/inferno-os/source/browse/utils/5c/sgen.c
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
		p-&gt;to.type = D_REG;
		p-&gt;to.reg = REGRET;
	}
	if(n &amp; 2) {
		gins(ANOP, Z, Z);
		p-&gt;to.type = D_FREG;
		p-&gt;to.reg = FREGRET;
	}
}

/*
 *	calculate addressability as follows
 *		CONST ==&gt; 20		$value
 *		NAME ==&gt; 10		name
 *		REGISTER ==&gt; 11		register
 *		INDREG ==&gt; 12		*[(reg)+offset]
 *		&amp;10 ==&gt; 2		$name
 *		ADD(2, 20) ==&gt; 2	$name+offset
 *		ADD(3, 20) ==&gt; 3	$(reg)+offset
 *		&amp;12 ==&gt; 3		$(reg)+offset
 *		*11 ==&gt; 11		??
 *		*2 ==&gt; 10		name
 *		*3 ==&gt; 12		*(reg)+offset
 *	calculate complexity (number of registers)
 */
void
xcom(Node *n)
{
	Node *l, *r;
	int t;

	if(n == Z)
		return;
	l = n-&gt;left;
	r = n-&gt;right;
	n-&gt;addable = 0;
	n-&gt;complex = 0;
	switch(n-&gt;op) {
	case OCONST:
		n-&gt;addable = 20;
		return;

	case OREGISTER:
		n-&gt;addable = 11;
		return;

	case OINDREG:
		n-&gt;addable = 12;
		return;

	case ONAME:
		n-&gt;addable = 10;
		return;

	case OADDR:
		xcom(l);
		if(l-&gt;addable == 10)
			n-&gt;addable = 2;
		if(l-&gt;addable == 12)
			n-&gt;addable = 3;
		break;

	case OIND:
		xcom(l);
		if(l-&gt;addable == 11)
			n-&gt;addable = 12;
		if(l-&gt;addable == 3)
			n-&gt;addable = 12;
		if(l-&gt;addable == 2)
			n-&gt;addable = 10;
		break;

	case OADD:
		xcom(l);
		xcom(r);
		if(l-&gt;addable == 20) {
			if(r-&gt;addable == 2)
				n-&gt;addable = 2;
			if(r-&gt;addable == 3)
				n-&gt;addable = 3;
		}
		if(r-&gt;addable == 20) {
			if(l-&gt;addable == 2)
				n-&gt;addable = 2;
			if(l-&gt;addable == 3)
				n-&gt;addable = 3;
		}
		break;

	case OASLMUL:
	case OASMUL:
		xcom(l);
		xcom(r);
		t = vlog(r);
		if(t &gt;= 0) {
			n-&gt;op = OASASHL;
			r-&gt;vconst = t;
			r-&gt;type = types[TINT];
		}
		break;

	case OMUL:
	case OLMUL:
		xcom(l);
		xcom(r);
		t = vlog(r);
		if(t &gt;= 0) {
			n-&gt;op = OASHL;
			r-&gt;vconst = t;
			r-&gt;type = types[TINT];
		}
		t = vlog(l);
		if(t &gt;= 0) {
			n-&gt;op = OASHL;
			n-&gt;left = r;
			n-&gt;right = l;
			r = l;
			l = n-&gt;left;
			r-&gt;vconst = t;
			r-&gt;type = types[TINT];
		}
		break;

	case OASLDIV:
		xcom(l);
		xcom(r);
		t = vlog(r);
		if(t &gt;= 0) {
			n-&gt;op = OASLSHR;
			r-&gt;vconst = t;
			r-&gt;type = types[TINT];
		}
		break;

	case OLDIV:
		xcom(l);
		xcom(r);
		t = vlog(r);
		if(t &gt;= 0) {
			n-&gt;op = OLSHR;
			r-&gt;vconst = t;
			r-&gt;type = types[TINT];
		}
		break;

	case OASLMOD:
		xcom(l);
		xcom(r);
		t = vlog(r);
		if(t &gt;= 0) {
			n-&gt;op = OASAND;
			r-&gt;vconst--;
		}
		break;

	case OLMOD:
		xcom(l);
		xcom(r);
		t = vlog(r);
		if(t &gt;= 0) {
			n-&gt;op = OAND;
			r-&gt;vconst--;
		}
		break;

	default:
		if(l != Z)
			xcom(l);
		if(r != Z)
			xcom(r);
		break;
	}
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

	case OADD:
	case OXOR:
	case OAND:
	case OOR:
	case OEQ:
	case ONE:
		/*
		 * immediate operators, make const on right
		 */
		if(l-&gt;op == OCONST) {
			n-&gt;left = r;
			n-&gt;right = l;
		}
		break;
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

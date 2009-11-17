<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5g/list.c</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5g/list.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Inferno utils/5c/list.c
// http://code.google.com/p/inferno-os/source/browse/utils/5c/list.c
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

#include &#34;gg.h&#34;

// TODO(kaib): make 5g/list.c congruent with 5l/list.c

static	int	sconsize;
void
listinit(void)
{

	fmtinstall(&#39;A&#39;, Aconv);		// as
	fmtinstall(&#39;C&#39;, Cconv);		// conditional execution bit
	fmtinstall(&#39;P&#39;, Pconv);			// Prog*
	fmtinstall(&#39;D&#39;, Dconv);		// Addr*
	fmtinstall(&#39;Y&#39;, Yconv);		// sconst
	fmtinstall(&#39;R&#39;, Rconv);		// register
	fmtinstall(&#39;M&#39;, Mconv);		// names
}

int
Pconv(Fmt *fp)
{
	char str[STRINGSZ];
	Prog *p;

	p = va_arg(fp-&gt;args, Prog*);
	sconsize = 8;
	switch(p-&gt;as) {
	default:
		if(p-&gt;reg == NREG)
			snprint(str, sizeof(str), &#34;%.4ld (%L) %-7A%C	%D,%D&#34;, 
				p-&gt;loc, p-&gt;lineno, p-&gt;as, p-&gt;scond, &amp;p-&gt;from, &amp;p-&gt;to);
		else if (p-&gt;from.type != D_FREG)
			snprint(str, sizeof(str), &#34;%.4ld (%L) %-7A%C	%D,R%d,%D&#34;, 
				p-&gt;loc, p-&gt;lineno, p-&gt;as, p-&gt;scond, &amp;p-&gt;from, p-&gt;reg, &amp;p-&gt;to);
		else
			snprint(str, sizeof(str), &#34;%.4ld (%L) %-7A%C	%D,F%d,%D&#34;,
				p-&gt;loc, p-&gt;lineno, p-&gt;as, p-&gt;scond, &amp;p-&gt;from, p-&gt;reg, &amp;p-&gt;to);
		break;

	case ADATA:
	case AINIT:
	case ADYNT:
		snprint(str, sizeof(str), &#34;%.4ld (%L) %-7A	%D/%d,%D&#34;,
			p-&gt;loc, p-&gt;lineno, p-&gt;as, &amp;p-&gt;from, p-&gt;reg, &amp;p-&gt;to);
		break;
	}
	return fmtstrcpy(fp, str);
}

int
Dconv(Fmt *fp)
{
	char str[STRINGSZ];
	char *op;
	Addr *a;
	int i;
	int32 v;

	a = va_arg(fp-&gt;args, Addr*);
	i = a-&gt;type;
	switch(i) {

	default:
		sprint(str, &#34;GOK-type(%d)&#34;, a-&gt;type);
		break;

	case D_NONE:
		str[0] = 0;
		if(a-&gt;name != D_NONE || a-&gt;reg != NREG || a-&gt;sym != S)
			sprint(str, &#34;%M(R%d)(NONE)&#34;, a, a-&gt;reg);
		break;

	case D_CONST:
		if(a-&gt;reg != NREG)
			sprint(str, &#34;$%M(R%d)&#34;, a, a-&gt;reg);
		else
			sprint(str, &#34;$%M&#34;, a);
		break;

	case D_CONST2:
		sprint(str, &#34;$%d-%d&#34;, a-&gt;offset, a-&gt;offset2);
		break;

	case D_SHIFT:
		v = a-&gt;offset;
		op = &#34;&lt;&lt;&gt;&gt;-&gt;@&gt;&#34; + (((v&gt;&gt;5) &amp; 3) &lt;&lt; 1);
		if(v &amp; (1&lt;&lt;4))
			sprint(str, &#34;R%d%c%cR%d&#34;, v&amp;15, op[0], op[1], (v&gt;&gt;8)&amp;15);
		else
			sprint(str, &#34;R%d%c%c%d&#34;, v&amp;15, op[0], op[1], (v&gt;&gt;7)&amp;31);
		if(a-&gt;reg != NREG)
			sprint(str+strlen(str), &#34;(R%d)&#34;, a-&gt;reg);
		break;

	case D_OCONST:
		sprint(str, &#34;$*$%M&#34;, a);
		if(a-&gt;reg != NREG)
			sprint(str, &#34;%M(R%d)(CONST)&#34;, a, a-&gt;reg);
		break;

	case D_OREG:
		if(a-&gt;reg != NREG)
			sprint(str, &#34;%M(R%d)&#34;, a, a-&gt;reg);
		else
			sprint(str, &#34;%M&#34;, a);
		break;

	case D_REG:
		sprint(str, &#34;R%d&#34;, a-&gt;reg);
		if(a-&gt;name != D_NONE || a-&gt;sym != S)
			sprint(str, &#34;%M(R%d)(REG)&#34;, a, a-&gt;reg);
		break;

	case D_REGREG:
		sprint(str, &#34;(R%d,R%d)&#34;, a-&gt;reg, (int)a-&gt;offset);
		if(a-&gt;name != D_NONE || a-&gt;sym != S)
			sprint(str, &#34;%M(R%d)(REG)&#34;, a, a-&gt;reg);
		break;

	case D_FREG:
		sprint(str, &#34;F%d&#34;, a-&gt;reg);
		if(a-&gt;name != D_NONE || a-&gt;sym != S)
			sprint(str, &#34;%M(R%d)(REG)&#34;, a, a-&gt;reg);
		break;

	case D_BRANCH:
		if(a-&gt;sym != S)
			sprint(str, &#34;%s+%d(APC)&#34;, a-&gt;sym-&gt;name, a-&gt;offset);
		else
			sprint(str, &#34;%d(APC)&#34;, a-&gt;offset);
		break;

	case D_FCONST:
		snprint(str, sizeof(str), &#34;$(%.17e)&#34;, a-&gt;dval);
		break;

	case D_SCONST:
		snprint(str, sizeof(str), &#34;$\&#34;%Y\&#34;&#34;, a-&gt;sval);
		break;

		// TODO(kaib): Add back
//	case D_ADDR:
//		a-&gt;type = a-&gt;index;
//		a-&gt;index = D_NONE;
//		snprint(str, sizeof(str), &#34;$%D&#34;, a);
//		a-&gt;index = a-&gt;type;
//		a-&gt;type = D_ADDR;
//		goto conv;
	}
//conv:
	return fmtstrcpy(fp, str);
}

int
Aconv(Fmt *fp)
{
	int i;

	i = va_arg(fp-&gt;args, int);
	return fmtstrcpy(fp, anames[i]);
}

char*	strcond[16] =
{
	&#34;.EQ&#34;,
	&#34;.NE&#34;,
	&#34;.HS&#34;,
	&#34;.LO&#34;,
	&#34;.MI&#34;,
	&#34;.PL&#34;,
	&#34;.VS&#34;,
	&#34;.VC&#34;,
	&#34;.HI&#34;,
	&#34;.LS&#34;,
	&#34;.GE&#34;,
	&#34;.LT&#34;,
	&#34;.GT&#34;,
	&#34;.LE&#34;,
	&#34;&#34;,
	&#34;.NV&#34;
};

int
Cconv(Fmt *fp)
{
	char s[20];
	int c;

	c = va_arg(fp-&gt;args, int);
	strcpy(s, strcond[c &amp; C_SCOND]);
	if(c &amp; C_SBIT)
		strcat(s, &#34;.S&#34;);
	if(c &amp; C_PBIT)
		strcat(s, &#34;.P&#34;);
	if(c &amp; C_WBIT)
		strcat(s, &#34;.W&#34;);
	if(c &amp; C_UBIT)		/* ambiguous with FBIT */
		strcat(s, &#34;.U&#34;);
	return fmtstrcpy(fp, s);
}

int
Yconv(Fmt *fp)
{
	int i, c;
	char str[30], *p, *a;

	a = va_arg(fp-&gt;args, char*);
	p = str;
	for(i=0; i&lt;sconsize; i++) {
		c = a[i] &amp; 0xff;
		if((c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;z&#39;) ||
		   (c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;Z&#39;) ||
		   (c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;)) {
			*p++ = c;
			continue;
		}
		*p++ = &#39;\\&#39;;
		switch(c) {
		default:
			if(c &lt; 040 || c &gt;= 0177)
				break;	/* not portable */
			p[-1] = c;
			continue;
		case 0:
			*p++ = &#39;z&#39;;
			continue;
		case &#39;\\&#39;:
		case &#39;&#34;&#39;:
			*p++ = c;
			continue;
		case &#39;\n&#39;:
			*p++ = &#39;n&#39;;
			continue;
		case &#39;\t&#39;:
			*p++ = &#39;t&#39;;
			continue;
		}
		*p++ = (c&gt;&gt;6) + &#39;0&#39;;
		*p++ = ((c&gt;&gt;3) &amp; 7) + &#39;0&#39;;
		*p++ = (c &amp; 7) + &#39;0&#39;;
	}
	*p = 0;
	return fmtstrcpy(fp, str);
}

int
Rconv(Fmt *fp)
{
	int r;
	char str[30];

	r = va_arg(fp-&gt;args, int);
	snprint(str, sizeof(str), &#34;R%d&#34;, r);
	return fmtstrcpy(fp, str);
}

int
Mconv(Fmt *fp)
{
	char str[STRINGSZ];
	Addr *a;

	a = va_arg(fp-&gt;args, Addr*);
	switch(a-&gt;name) {
	default:
		snprint(str, sizeof(str),  &#34;GOK-name(%d)&#34;, a-&gt;name);
		break;

	case D_NONE:
		snprint(str, sizeof(str), &#34;%d&#34;, a-&gt;offset);
		break;

	case D_EXTERN:
		snprint(str, sizeof(str), &#34;%S+%d(SB)&#34;, a-&gt;sym, a-&gt;offset);
		break;

	case D_STATIC:
		snprint(str, sizeof(str), &#34;%S&lt;&gt;+%d(SB)&#34;, a-&gt;sym, a-&gt;offset);
		break;

	case D_AUTO:
		snprint(str, sizeof(str), &#34;%S+%d(SP)&#34;, a-&gt;sym, a-&gt;offset);
		break;

	case D_PARAM:
		snprint(str, sizeof(str), &#34;%S+%d(FP)&#34;, a-&gt;sym, a-&gt;offset);
		break;
	}
	return fmtstrcpy(fp, str);
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6l/list.c</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/6l/list.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/6l/list.c
// http://code.google.com/p/inferno-os/source/browse/utils/6l/list.c
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

#include	&#34;l.h&#34;
#include	&#34;../ld/lib.h&#34;

static	Prog*	bigP;

void
listinit(void)
{

	fmtinstall(&#39;R&#39;, Rconv);
	fmtinstall(&#39;A&#39;, Aconv);
	fmtinstall(&#39;D&#39;, Dconv);
	fmtinstall(&#39;S&#39;, Sconv);
	fmtinstall(&#39;P&#39;, Pconv);
}

int
Pconv(Fmt *fp)
{
	char str[STRINGSZ], str1[STRINGSZ];
	Prog *p;

	p = va_arg(fp-&gt;args, Prog*);
	if(p == P)
		return fmtstrcpy(fp, &#34;&lt;P&gt;&#34;);

	bigP = p;

	snprint(str1, sizeof(str1), &#34;(%ld)&#34;, p-&gt;line);
	switch(p-&gt;as) {
	case ATEXT:
		if(p-&gt;from.scale) {
			snprint(str, sizeof(str), &#34;%-7s %-7A %D,%d,%lD&#34;,
				str1, p-&gt;as, &amp;p-&gt;from, p-&gt;from.scale, &amp;p-&gt;to);
			break;
		}
		snprint(str, sizeof(str), &#34;%-7s %-7A %D,%lD&#34;,
			str1, p-&gt;as, &amp;p-&gt;from, &amp;p-&gt;to);
		break;

	default:
		snprint(str, sizeof(str), &#34;%-7s %-7A %D,%D&#34;,
			str1, p-&gt;as, &amp;p-&gt;from, &amp;p-&gt;to);
		break;

	case ADATA:
	case AINIT:
	case ADYNT:
		snprint(str, sizeof(str), &#34;%-7s %-7A %D/%d,%D&#34;,
			str1, p-&gt;as, &amp;p-&gt;from, p-&gt;from.scale, &amp;p-&gt;to);
		break;
	}
	bigP = P;
	return fmtstrcpy(fp, str);
}

int
Aconv(Fmt *fp)
{
	int i;

	i = va_arg(fp-&gt;args, int);
	return fmtstrcpy(fp, anames[i]);
}

int
Dconv(Fmt *fp)
{
	char str[40], s[20];
	Adr *a;
	int i;

	a = va_arg(fp-&gt;args, Adr*);
	i = a-&gt;type;

	if(fp-&gt;flags &amp; FmtLong) {
		if(i != D_CONST) {
			// ATEXT dst is not constant
			snprint(str, sizeof(str), &#34;!!%D&#34;, a);
			goto brk;
		}
		parsetextconst(a-&gt;offset);
		if(textarg == 0) {
			snprint(str, sizeof(str), &#34;$%lld&#34;, textstksiz);
			goto brk;
		}
		snprint(str, sizeof(str), &#34;$%lld-%lld&#34;, textstksiz, textarg);
		goto brk;
	}

	if(i &gt;= D_INDIR) {
		if(a-&gt;offset)
			snprint(str, sizeof(str), &#34;%lld(%R)&#34;, a-&gt;offset, i-D_INDIR);
		else
			snprint(str, sizeof(str), &#34;(%R)&#34;, i-D_INDIR);
		goto brk;
	}
	switch(i) {

	default:
		if(a-&gt;offset)
			snprint(str, sizeof(str), &#34;$%lld,%R&#34;, a-&gt;offset, i);
		else
			snprint(str, sizeof(str), &#34;%R&#34;, i);
		break;

	case D_NONE:
		str[0] = 0;
		break;

	case D_BRANCH:
		if(bigP != P &amp;&amp; bigP-&gt;pcond != P)
			if(a-&gt;sym != S)
				snprint(str, sizeof(str), &#34;%llux+%s&#34;, bigP-&gt;pcond-&gt;pc,
					a-&gt;sym-&gt;name);
			else
				snprint(str, sizeof(str), &#34;%llux&#34;, bigP-&gt;pcond-&gt;pc);
		else
			snprint(str, sizeof(str), &#34;%lld(PC)&#34;, a-&gt;offset);
		break;

	case D_EXTERN:
		if(a-&gt;sym) {
			snprint(str, sizeof(str), &#34;%s+%lld(SB)&#34;, a-&gt;sym-&gt;name, a-&gt;offset);
			break;
		}
		snprint(str, sizeof(str), &#34;!!noname!!+%lld(SB)&#34;, a-&gt;offset);
		break;

	case D_STATIC:
		if(a-&gt;sym) {
			snprint(str, sizeof(str), &#34;%s&lt;%d&gt;+%lld(SB)&#34;, a-&gt;sym-&gt;name,
				a-&gt;sym-&gt;version, a-&gt;offset);
			break;
		}
		snprint(str, sizeof(str), &#34;!!noname!!&lt;999&gt;+%lld(SB)&#34;, a-&gt;offset);
		break;

	case D_AUTO:
		if(a-&gt;sym) {
			snprint(str, sizeof(str), &#34;%s+%lld(SP)&#34;, a-&gt;sym-&gt;name, a-&gt;offset);
			break;
		}
		snprint(str, sizeof(str), &#34;!!noname!!+%lld(SP)&#34;, a-&gt;offset);
		break;

	case D_PARAM:
		if(a-&gt;sym) {
			snprint(str, sizeof(str), &#34;%s+%lld(%s)&#34;, a-&gt;sym-&gt;name, a-&gt;offset, paramspace);
			break;
		}
		snprint(str, sizeof(str), &#34;!!noname!!+%lld(%s)&#34;, a-&gt;offset, paramspace);
		break;

	case D_CONST:
		snprint(str, sizeof(str), &#34;$%lld&#34;, a-&gt;offset);
		break;

	case D_FCONST:
		snprint(str, sizeof(str), &#34;$(%.8lux,%.8lux)&#34;, a-&gt;ieee.h, a-&gt;ieee.l);
		break;

	case D_SCONST:
		snprint(str, sizeof(str), &#34;$\&#34;%S\&#34;&#34;, a-&gt;scon);
		break;

	case D_ADDR:
		a-&gt;type = a-&gt;index;
		a-&gt;index = D_NONE;
		snprint(str, sizeof(str), &#34;$%D&#34;, a);
		a-&gt;index = a-&gt;type;
		a-&gt;type = D_ADDR;
		goto conv;
	}
brk:
	if(a-&gt;index != D_NONE) {
		snprint(s, sizeof(s), &#34;(%R*%d)&#34;, a-&gt;index, a-&gt;scale);
		strcat(str, s);
	}
conv:
	fmtstrcpy(fp, str);
//	if(a-&gt;gotype)
//		fmtprint(fp, &#34;«%s»&#34;, a-&gt;gotype-&gt;name);
	return 0;

}

char*	regstr[] =
{
	&#34;AL&#34;,		/* [D_AL] */
	&#34;CL&#34;,
	&#34;DL&#34;,
	&#34;BL&#34;,
	&#34;SPB&#34;,
	&#34;BPB&#34;,
	&#34;SIB&#34;,
	&#34;DIB&#34;,
	&#34;R8B&#34;,
	&#34;R9B&#34;,
	&#34;R10B&#34;,
	&#34;R11B&#34;,
	&#34;R12B&#34;,
	&#34;R13B&#34;,
	&#34;R14B&#34;,
	&#34;R15B&#34;,

	&#34;AX&#34;,		/* [D_AX] */
	&#34;CX&#34;,
	&#34;DX&#34;,
	&#34;BX&#34;,
	&#34;SP&#34;,
	&#34;BP&#34;,
	&#34;SI&#34;,
	&#34;DI&#34;,
	&#34;R8&#34;,
	&#34;R9&#34;,
	&#34;R10&#34;,
	&#34;R11&#34;,
	&#34;R12&#34;,
	&#34;R13&#34;,
	&#34;R14&#34;,
	&#34;R15&#34;,

	&#34;AH&#34;,
	&#34;CH&#34;,
	&#34;DH&#34;,
	&#34;BH&#34;,

	&#34;F0&#34;,		/* [D_F0] */
	&#34;F1&#34;,
	&#34;F2&#34;,
	&#34;F3&#34;,
	&#34;F4&#34;,
	&#34;F5&#34;,
	&#34;F6&#34;,
	&#34;F7&#34;,

	&#34;M0&#34;,
	&#34;M1&#34;,
	&#34;M2&#34;,
	&#34;M3&#34;,
	&#34;M4&#34;,
	&#34;M5&#34;,
	&#34;M6&#34;,
	&#34;M7&#34;,

	&#34;X0&#34;,
	&#34;X1&#34;,
	&#34;X2&#34;,
	&#34;X3&#34;,
	&#34;X4&#34;,
	&#34;X5&#34;,
	&#34;X6&#34;,
	&#34;X7&#34;,
	&#34;X8&#34;,
	&#34;X9&#34;,
	&#34;X10&#34;,
	&#34;X11&#34;,
	&#34;X12&#34;,
	&#34;X13&#34;,
	&#34;X14&#34;,
	&#34;X15&#34;,

	&#34;CS&#34;,		/* [D_CS] */
	&#34;SS&#34;,
	&#34;DS&#34;,
	&#34;ES&#34;,
	&#34;FS&#34;,
	&#34;GS&#34;,

	&#34;GDTR&#34;,		/* [D_GDTR] */
	&#34;IDTR&#34;,		/* [D_IDTR] */
	&#34;LDTR&#34;,		/* [D_LDTR] */
	&#34;MSW&#34;,		/* [D_MSW] */
	&#34;TASK&#34;,		/* [D_TASK] */

	&#34;CR0&#34;,		/* [D_CR] */
	&#34;CR1&#34;,
	&#34;CR2&#34;,
	&#34;CR3&#34;,
	&#34;CR4&#34;,
	&#34;CR5&#34;,
	&#34;CR6&#34;,
	&#34;CR7&#34;,
	&#34;CR8&#34;,
	&#34;CR9&#34;,
	&#34;CR10&#34;,
	&#34;CR11&#34;,
	&#34;CR12&#34;,
	&#34;CR13&#34;,
	&#34;CR14&#34;,
	&#34;CR15&#34;,

	&#34;DR0&#34;,		/* [D_DR] */
	&#34;DR1&#34;,
	&#34;DR2&#34;,
	&#34;DR3&#34;,
	&#34;DR4&#34;,
	&#34;DR5&#34;,
	&#34;DR6&#34;,
	&#34;DR7&#34;,

	&#34;TR0&#34;,		/* [D_TR] */
	&#34;TR1&#34;,
	&#34;TR2&#34;,
	&#34;TR3&#34;,
	&#34;TR4&#34;,
	&#34;TR5&#34;,
	&#34;TR6&#34;,
	&#34;TR7&#34;,

	&#34;NONE&#34;,		/* [D_NONE] */
};

int
Rconv(Fmt *fp)
{
	char str[20];
	int r;

	r = va_arg(fp-&gt;args, int);
	if(r &gt;= D_AL &amp;&amp; r &lt;= D_NONE)
		snprint(str, sizeof(str), &#34;%s&#34;, regstr[r-D_AL]);
	else
		snprint(str, sizeof(str), &#34;gok(%d)&#34;, r);

	return fmtstrcpy(fp, str);
}

int
Sconv(Fmt *fp)
{
	int i, c;
	char str[30], *p, *a;

	a = va_arg(fp-&gt;args, char*);
	p = str;
	for(i=0; i&lt;sizeof(double); i++) {
		c = a[i] &amp; 0xff;
		if(c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;z&#39; ||
		   c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;Z&#39; ||
		   c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;) {
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

void
diag(char *fmt, ...)
{
	char buf[STRINGSZ], *tn;
	va_list arg;

	tn = &#34;??none??&#34;;
	if(curtext != P &amp;&amp; curtext-&gt;from.sym != S)
		tn = curtext-&gt;from.sym-&gt;name;
	va_start(arg, fmt);
	vseprint(buf, buf+sizeof(buf), fmt, arg);
	va_end(arg);
	print(&#34;%s: %s\n&#34;, tn, buf);

	nerrors++;
	if(nerrors &gt; 20) {
		print(&#34;too many errors\n&#34;);
		errorexit();
	}
}

void
parsetextconst(vlong arg)
{
	textstksiz = arg &amp; 0xffffffffLL;
	if(textstksiz &amp; 0x80000000LL)
		textstksiz = -(-textstksiz &amp; 0xffffffffLL);

	textarg = (arg &gt;&gt; 32) &amp; 0xffffffffLL;
	if(textarg &amp; 0x80000000LL)
		textarg = 0;
	textarg = (textarg+7) &amp; ~7LL;
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

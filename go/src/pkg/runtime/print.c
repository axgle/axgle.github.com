<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/print.c</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/print.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;

//static Lock debuglock;

void
dump(byte *p, int32 n)
{
	int32 i;

	for(i=0; i&lt;n; i++) {
		runtime·printpointer((byte*)(p[i]&gt;&gt;4));
		runtime·printpointer((byte*)(p[i]&amp;0xf));
		if((i&amp;15) == 15)
			prints(&#34;\n&#34;);
		else
			prints(&#34; &#34;);
	}
	if(n &amp; 15)
		prints(&#34;\n&#34;);
}

void
prints(int8 *s)
{
	write(fd, s, findnull((byte*)s));
}

// Very simple printf.  Only for debugging prints.
// Do not add to this without checking with Rob.
void
printf(int8 *s, ...)
{
	int8 *p, *lp;
	byte *arg, *narg;

//	lock(&amp;debuglock);

	lp = p = s;
	arg = (byte*)(&amp;s+1);
	for(; *p; p++) {
		if(*p != &#39;%&#39;)
			continue;
		if(p &gt; lp)
			write(fd, lp, p-lp);
		p++;
		narg = nil;
		switch(*p) {
		case &#39;d&#39;:	// 32-bit
		case &#39;x&#39;:
			narg = arg + 4;
			break;
		case &#39;D&#39;:	// 64-bit
		case &#39;X&#39;:
			if(sizeof(uintptr) == 8 &amp;&amp; ((uint32)(uint64)arg)&amp;4)
				arg += 4;
			narg = arg + 8;
			break;
		case &#39;p&#39;:	// pointer-sized
		case &#39;s&#39;:
			if(sizeof(uintptr) == 8 &amp;&amp; ((uint32)(uint64)arg)&amp;4)
				arg += 4;
			narg = arg + sizeof(uintptr);
			break;
		case &#39;S&#39;:	// pointer-aligned but bigger
			if(sizeof(uintptr) == 8 &amp;&amp; ((uint32)(uint64)arg)&amp;4)
				arg += 4;
			narg = arg + sizeof(String);
			break;
		}
		switch(*p) {
		case &#39;d&#39;:
			runtime·printint(*(int32*)arg);
			break;
		case &#39;D&#39;:
			runtime·printint(*(int64*)arg);
			break;
		case &#39;x&#39;:
			runtime·printhex(*(uint32*)arg);
			break;
		case &#39;X&#39;:
			runtime·printhex(*(uint64*)arg);
			break;
		case &#39;p&#39;:
			runtime·printpointer(*(void**)arg);
			break;
		case &#39;s&#39;:
			prints(*(int8**)arg);
			break;
		case &#39;S&#39;:
			runtime·printstring(*(String*)arg);
			break;
		}
		arg = narg;
		lp = p+1;
	}
	if(p &gt; lp)
		write(fd, lp, p-lp);

//	unlock(&amp;debuglock);
}


void
runtime·printpc(void *p)
{
	prints(&#34;PC=&#34;);
	runtime·printhex((uint64)runtime·getcallerpc(p));
}

void
runtime·printbool(bool v)
{
	if(v) {
		write(fd, (byte*)&#34;true&#34;, 4);
		return;
	}
	write(fd, (byte*)&#34;false&#34;, 5);
}

void
runtime·printfloat(float64 v)
{
	byte buf[20];
	int32 e, s, i, n;
	float64 h;

	if(isNaN(v)) {
		write(fd, &#34;NaN&#34;, 3);
		return;
	}
	if(isInf(v, 0)) {
		write(fd, &#34;+Inf&#34;, 4);
		return;
	}
	if(isInf(v, -1)) {
		write(fd, &#34;+Inf&#34;, 4);
		return;
	}


	n = 7;	// digits printed
	e = 0;	// exp
	s = 0;	// sign
	if(v != 0) {
		// sign
		if(v &lt; 0) {
			v = -v;
			s = 1;
		}

		// normalize
		while(v &gt;= 10) {
			e++;
			v /= 10;
		}
		while(v &lt; 1) {
			e--;
			v *= 10;
		}

		// round
		h = 5;
		for(i=0; i&lt;n; i++)
			h /= 10;
		v += h;
		if(v &gt;= 10) {
			e++;
			v /= 10;
		}
	}

	// format +d.dddd+edd
	buf[0] = &#39;+&#39;;
	if(s)
		buf[0] = &#39;-&#39;;
	for(i=0; i&lt;n; i++) {
		s = v;
		buf[i+2] = s+&#39;0&#39;;
		v -= s;
		v *= 10.;
	}
	buf[1] = buf[2];
	buf[2] = &#39;.&#39;;

	buf[n+2] = &#39;e&#39;;
	buf[n+3] = &#39;+&#39;;
	if(e &lt; 0) {
		e = -e;
		buf[n+3] = &#39;-&#39;;
	}

	buf[n+4] = (e/100) + &#39;0&#39;;
	buf[n+5] = (e/10)%10 + &#39;0&#39;;
	buf[n+6] = (e%10) + &#39;0&#39;;
	write(fd, buf, n+7);
}

void
runtime·printuint(uint64 v)
{
	byte buf[100];
	int32 i;

	for(i=nelem(buf)-1; i&gt;0; i--) {
		buf[i] = v%10 + &#39;0&#39;;
		if(v &lt; 10)
			break;
		v = v/10;
	}
	write(fd, buf+i, nelem(buf)-i);
}

void
runtime·printint(int64 v)
{
	if(v &lt; 0) {
		write(fd, &#34;-&#34;, 1);
		v = -v;
	}
	runtime·printuint(v);
}

void
runtime·printhex(uint64 v)
{
	static int8 *dig = &#34;0123456789abcdef&#34;;
	byte buf[100];
	int32 i;

	i=nelem(buf);
	for(; v&gt;0; v/=16)
		buf[--i] = dig[v%16];
	if(i == nelem(buf))
		buf[--i] = &#39;0&#39;;
	buf[--i] = &#39;x&#39;;
	buf[--i] = &#39;0&#39;;
	write(fd, buf+i, nelem(buf)-i);
}

void
runtime·printpointer(void *p)
{
	runtime·printhex((uint64)p);
}

void
runtime·printstring(String v)
{
	extern int32 maxstring;

	if(v.len &gt; maxstring) {
		write(fd, &#34;[invalid string]&#34;, 16);
		return;
	}
	if(v.len &gt; 0)
		write(fd, v.str, v.len);
}

void
runtime·printsp(void)
{
	write(fd, &#34; &#34;, 1);
}

void
runtime·printnl(void)
{
	write(fd, &#34;\n&#34;, 1);
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

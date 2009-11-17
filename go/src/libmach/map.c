<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/map.c</title>

  <link rel="stylesheet" type="text/css" href="../../doc/style.css">
  <script type="text/javascript" src="../../doc/godocs.js"></script>

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
        <a href="../../index.html"><img src="../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../doc/install.html">Install Go</a></li>
    <li><a href="../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Text file src/libmach/map.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Inferno libmach/map.c and
// Plan 9 from User Space src/libmach/map.c
//
// http://code.swtch.com/plan9port/src/tip/src/libmach/map.c
// http://code.google.com/p/inferno-os/source/browse/utils/libmach/map.c
//
//
//	Copyright © 1994-1999 Lucent Technologies Inc.
//	Power PC support Copyright © 1995-2004 C H Forsyth (forsyth@terzarima.net).
//	Portions Copyright © 1997-1999 Vita Nuova Limited.
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).
//	Revisions Copyright © 2000-2004 Lucent Technologies Inc. and others.
//	Portions Copyright © 2001-2007 Russ Cox.
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

/*
 * file map routines
 */
#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;mach.h&gt;

Map *
newmap(Map *map, int n)
{
	int size;

	size = sizeof(Map)+(n-1)*sizeof(Seg);
	if (map == 0)
		map = malloc(size);
	else
		map = realloc(map, size);
	if (map == 0) {
		werrstr(&#34;out of memory: %r&#34;);
		return 0;
	}
	memset(map, 0, size);
	map-&gt;nsegs = n;
	return map;
}

int
setmap(Map *map, int fd, uvlong b, uvlong e, vlong f, char *name, Maprw *rw)
{
	int i;

	if (map == 0)
		return 0;
	for (i = 0; i &lt; map-&gt;nsegs; i++)
		if (!map-&gt;seg[i].inuse)
			break;
	if (i &gt;= map-&gt;nsegs)
		return 0;
	map-&gt;seg[i].b = b;
	map-&gt;seg[i].e = e;
	map-&gt;seg[i].f = f;
	map-&gt;seg[i].inuse = 1;
	map-&gt;seg[i].name = name;
	map-&gt;seg[i].fd = fd;
	map-&gt;seg[i].rw = rw;
	return 1;
}

/*
static uvlong
stacktop(int pid)
{
	char buf[64];
	int fd;
	int n;
	char *cp;

	snprint(buf, sizeof(buf), &#34;/proc/%d/segment&#34;, pid);
	fd = open(buf, 0);
	if (fd &lt; 0)
		return 0;
	n = read(fd, buf, sizeof(buf)-1);
	close(fd);
	buf[n] = 0;
	if (strncmp(buf, &#34;Stack&#34;, 5))
		return 0;
	for (cp = buf+5; *cp &amp;&amp; *cp == &#39; &#39;; cp++)
		;
	if (!*cp)
		return 0;
	cp = strchr(cp, &#39; &#39;);
	if (!cp)
		return 0;
	while (*cp &amp;&amp; *cp == &#39; &#39;)
		cp++;
	if (!*cp)
		return 0;
	return strtoull(cp, 0, 16);
}
*/

int
findseg(Map *map, char *name)
{
	int i;

	if (!map)
		return -1;
	for (i = 0; i &lt; map-&gt;nsegs; i++)
		if (map-&gt;seg[i].inuse &amp;&amp; !strcmp(map-&gt;seg[i].name, name))
			return i;
	return -1;
}

void
unusemap(Map *map, int i)
{
	if (map != 0 &amp;&amp; 0 &lt;= i &amp;&amp; i &lt; map-&gt;nsegs)
		map-&gt;seg[i].inuse = 0;
}

int
fdrw(Map *map, Seg *s, uvlong addr, void *v, uint n, int isread)
{
	int tot, m;

	for(tot=0; tot&lt;n; tot+=m){
		if(isread)
			m = pread(s-&gt;fd, (uchar*)v+tot, n-tot, addr+tot);
		else
			m = pwrite(s-&gt;fd, (uchar*)v+tot, n-tot, addr+tot);
		if(m == 0){
			werrstr(&#34;short %s&#34;, isread ? &#34;read&#34; : &#34;write&#34;);
			return -1;
		}
		if(m &lt; 0){
			werrstr(&#34;%s %d at %#llux (+%#llux): %r&#34;, isread ? &#34;read&#34; : &#34;write&#34;, n, addr, s-&gt;f);
			return -1;
		}
	}
	return 0;
}


Map*
loadmap(Map *map, int fd, Fhdr *fp)
{
	map = newmap(map, 2);
	if (map == 0)
		return 0;

	map-&gt;seg[0].b = fp-&gt;txtaddr;
	map-&gt;seg[0].e = fp-&gt;txtaddr+fp-&gt;txtsz;
	map-&gt;seg[0].f = fp-&gt;txtoff;
	map-&gt;seg[0].fd = fd;
	map-&gt;seg[0].inuse = 1;
	map-&gt;seg[0].name = &#34;text&#34;;
	map-&gt;seg[0].rw = fdrw;
	map-&gt;seg[1].b = fp-&gt;dataddr;
	map-&gt;seg[1].e = fp-&gt;dataddr+fp-&gt;datsz;
	map-&gt;seg[1].f = fp-&gt;datoff;
	map-&gt;seg[1].fd = fd;
	map-&gt;seg[1].inuse = 1;
	map-&gt;seg[1].name = &#34;data&#34;;
	map-&gt;seg[0].rw = fdrw;
	return map;
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

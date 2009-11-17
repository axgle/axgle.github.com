<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/_p9dir.c</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/lib9/_p9dir.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
/*
Plan 9 from User Space src/lib9/_p9dir.c
http://code.swtch.com/plan9port/src/tip/src/lib9/_p9dir.c

Copyright 2001-2007 Russ Cox.  All Rights Reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the &#34;Software&#34;), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

#include &lt;u.h&gt;
#define NOPLAN9DEFINES
#include &lt;libc.h&gt;
#include &lt;sys/types.h&gt;
#include &lt;sys/stat.h&gt;
#include &lt;dirent.h&gt;
#include &lt;pwd.h&gt;
#include &lt;grp.h&gt;

/*
 * No need for a real disk size function here:
 * the Go build isn&#39;t looking at raw disk devices,
 * so this avoids portability problems.
 */
#define  _HAVEDISKSIZE
static vlong
disksize(int fd, int x)
{
	return 0;
}

/*
 * Caching the last group and passwd looked up is
 * a significant win (stupidly enough) on most systems.
 * It&#39;s not safe for threaded programs, but neither is using
 * getpwnam in the first place, so I&#39;m not too worried.
 */
int
_p9dir(struct stat *lst, struct stat *st, char *name, Dir *d, char **str, char *estr)
{
	char *s;
	char tmp[20];
	static struct group *g;
	static struct passwd *p;
	static int gid, uid;
	int sz, fd;

	fd = -1;
	USED(fd);
	sz = 0;
	if(d)
		memset(d, 0, sizeof *d);

	/* name */
	s = strrchr(name, &#39;/&#39;);
	if(s)
		s++;
	if(!s || !*s)
		s = name;
	if(*s == &#39;/&#39;)
		s++;
	if(*s == 0)
		s = &#34;/&#34;;
	if(d){
		if(*str + strlen(s)+1 &gt; estr)
			d-&gt;name = &#34;oops&#34;;
		else{
			strcpy(*str, s);
			d-&gt;name = *str;
			*str += strlen(*str)+1;
		}
	}
	sz += strlen(s)+1;

	/* user */
	if(p == nil || st-&gt;st_uid != uid || p-&gt;pw_uid != uid){
		snprint(tmp, sizeof tmp, &#34;%d&#34;, (int)st-&gt;st_uid);
		s = tmp;
	}else
		s = p-&gt;pw_name;
	sz += strlen(s)+1;
	if(d){
		if(*str+strlen(s)+1 &gt; estr)
			d-&gt;uid = &#34;oops&#34;;
		else{
			strcpy(*str, s);
			d-&gt;uid = *str;
			*str += strlen(*str)+1;
		}
	}

	/* group */
	if(g == nil || st-&gt;st_gid != gid || g-&gt;gr_gid != gid){
		snprint(tmp, sizeof tmp, &#34;%d&#34;, (int)st-&gt;st_gid);
		s = tmp;
	}else
		s = g-&gt;gr_name;
	sz += strlen(s)+1;
	if(d){
		if(*str + strlen(s)+1 &gt; estr)
			d-&gt;gid = &#34;oops&#34;;
		else{
			strcpy(*str, s);
			d-&gt;gid = *str;
			*str += strlen(*str)+1;
		}
	}

	if(d){
		d-&gt;type = &#39;M&#39;;

		d-&gt;muid = &#34;&#34;;
		d-&gt;qid.path = ((uvlong)st-&gt;st_dev&lt;&lt;32) | st-&gt;st_ino;
#ifdef _HAVESTGEN
		d-&gt;qid.vers = st-&gt;st_gen;
#endif
		if(d-&gt;qid.vers == 0)
			d-&gt;qid.vers = st-&gt;st_mtime + st-&gt;st_ctime;
		d-&gt;mode = st-&gt;st_mode&amp;0777;
		d-&gt;atime = st-&gt;st_atime;
		d-&gt;mtime = st-&gt;st_mtime;
		d-&gt;length = st-&gt;st_size;

		if(S_ISDIR(st-&gt;st_mode)){
			d-&gt;length = 0;
			d-&gt;mode |= DMDIR;
			d-&gt;qid.type = QTDIR;
		}
		if(S_ISLNK(lst-&gt;st_mode))	/* yes, lst not st */
			d-&gt;mode |= DMSYMLINK;
		if(S_ISFIFO(st-&gt;st_mode))
			d-&gt;mode |= DMNAMEDPIPE;
		if(S_ISSOCK(st-&gt;st_mode))
			d-&gt;mode |= DMSOCKET;
		if(S_ISBLK(st-&gt;st_mode)){
			d-&gt;mode |= DMDEVICE;
			d-&gt;qid.path = (&#39;b&#39;&lt;&lt;16)|st-&gt;st_rdev;
		}
		if(S_ISCHR(st-&gt;st_mode)){
			d-&gt;mode |= DMDEVICE;
			d-&gt;qid.path = (&#39;c&#39;&lt;&lt;16)|st-&gt;st_rdev;
		}
		/* fetch real size for disks */
#ifdef _HAVEDISKSIZE
		if(S_ISBLK(st-&gt;st_mode) &amp;&amp; (fd = open(name, O_RDONLY)) &gt;= 0){
			d-&gt;length = disksize(fd, major(st-&gt;st_dev));
			close(fd);
		}
#endif
#if defined(DIOCGMEDIASIZE)
		if(isdisk(st)){
			int fd;
			off_t mediasize;

			if((fd = open(name, O_RDONLY)) &gt;= 0){
				if(ioctl(fd, DIOCGMEDIASIZE, &amp;mediasize) &gt;= 0)
					d-&gt;length = mediasize;
				close(fd);
			}
		}
#elif defined(_HAVEDISKLABEL)
		if(isdisk(st)){
			int fd, n;
			struct disklabel lab;

			if((fd = open(name, O_RDONLY)) &lt; 0)
				goto nosize;
			if(ioctl(fd, DIOCGDINFO, &amp;lab) &lt; 0)
				goto nosize;
			n = minor(st-&gt;st_rdev)&amp;7;
			if(n &gt;= lab.d_npartitions)
				goto nosize;

			d-&gt;length = (vlong)(lab.d_partitions[n].p_size) * lab.d_secsize;

		nosize:
			if(fd &gt;= 0)
				close(fd);
		}
#endif
	}

	return sz;
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

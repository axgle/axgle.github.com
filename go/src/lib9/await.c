<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/await.c</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/lib9/await.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
/*
Plan 9 from User Space src/lib9/await.c
http://code.swtch.com/plan9port/src/tip/src/lib9/await.c

Copyright 2001-2007 Russ Cox.  All Rights Reserved.
Portions Copyright 2009 The Go Authors.  All Rights Reserved.

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

#define NOPLAN9DEFINES
#include &lt;u.h&gt;
#include &lt;libc.h&gt;

#include &lt;signal.h&gt;
#include &lt;sys/types.h&gt;
#include &lt;sys/wait.h&gt;
#include &lt;sys/time.h&gt;
#include &lt;sys/resource.h&gt;

#ifndef WCOREDUMP	/* not on Mac OS X Tiger */
#define WCOREDUMP(status) 0
#endif

static struct {
	int sig;
	char *str;
} tab[] = {
	SIGHUP,		&#34;hangup&#34;,
	SIGINT,		&#34;interrupt&#34;,
	SIGQUIT,		&#34;quit&#34;,
	SIGILL,		&#34;sys: illegal instruction&#34;,
	SIGTRAP,		&#34;sys: breakpoint&#34;,
	SIGABRT,		&#34;sys: abort&#34;,
#ifdef SIGEMT
	SIGEMT,		&#34;sys: emulate instruction executed&#34;,
#endif
	SIGFPE,		&#34;sys: fp: trap&#34;,
	SIGKILL,		&#34;sys: kill&#34;,
	SIGBUS,		&#34;sys: bus error&#34;,
	SIGSEGV,		&#34;sys: segmentation violation&#34;,
	SIGALRM,		&#34;alarm&#34;,
	SIGTERM,		&#34;kill&#34;,
	SIGURG,		&#34;sys: urgent condition on socket&#34;,
	SIGSTOP,		&#34;sys: stop&#34;,
	SIGTSTP,		&#34;sys: tstp&#34;,
	SIGCONT,		&#34;sys: cont&#34;,
	SIGCHLD,		&#34;sys: child&#34;,
	SIGTTIN,		&#34;sys: ttin&#34;,
	SIGTTOU,		&#34;sys: ttou&#34;,
#ifdef SIGIO	/* not on Mac OS X Tiger */
	SIGIO,		&#34;sys: i/o possible on fd&#34;,
#endif
	SIGXCPU,		&#34;sys: cpu time limit exceeded&#34;,
	SIGXFSZ,		&#34;sys: file size limit exceeded&#34;,
	SIGVTALRM,	&#34;sys: virtual time alarm&#34;,
	SIGPROF,		&#34;sys: profiling timer alarm&#34;,
#ifdef SIGWINCH	/* not on Mac OS X Tiger */
	SIGWINCH,	&#34;sys: window size change&#34;,
#endif
#ifdef SIGINFO
	SIGINFO,		&#34;sys: status request&#34;,
#endif
	SIGUSR1,		&#34;sys: usr1&#34;,
	SIGUSR2,		&#34;sys: usr2&#34;,
	SIGPIPE,		&#34;sys: write on closed pipe&#34;,
};

char*
_p9sigstr(int sig, char *tmp)
{
	int i;

	for(i=0; i&lt;nelem(tab); i++)
		if(tab[i].sig == sig)
			return tab[i].str;
	if(tmp == nil)
		return nil;
	sprint(tmp, &#34;sys: signal %d&#34;, sig);
	return tmp;
}

int
_p9strsig(char *s)
{
	int i;

	for(i=0; i&lt;nelem(tab); i++)
		if(strcmp(s, tab[i].str) == 0)
			return tab[i].sig;
	return 0;
}

static Waitmsg*
_wait(int pid4, int opt)
{
	int pid, status, cd;
	struct rusage ru;
	char tmp[64];
	ulong u, s;
	Waitmsg *w;

	w = malloc(sizeof *w + 200);
	if(w == nil)
		return nil;
	memset(w, 0, sizeof *w);
	w-&gt;msg = (char*)&amp;w[1];

	for(;;){
		/* On Linux, pid==-1 means anyone; on SunOS, it&#39;s pid==0. */
		if(pid4 == -1)
			pid = wait3(&amp;status, opt, &amp;ru);
		else
			pid = wait4(pid4, &amp;status, opt, &amp;ru);
		if(pid &lt;= 0) {
			free(w);
			return nil;
		}
		u = ru.ru_utime.tv_sec*1000+((ru.ru_utime.tv_usec+500)/1000);
		s = ru.ru_stime.tv_sec*1000+((ru.ru_stime.tv_usec+500)/1000);
		w-&gt;pid = pid;
		w-&gt;time[0] = u;
		w-&gt;time[1] = s;
		w-&gt;time[2] = u+s;
		if(WIFEXITED(status)){
			if(status)
				sprint(w-&gt;msg, &#34;%d&#34;, status);
			return w;
		}
		if(WIFSIGNALED(status)){
			cd = WCOREDUMP(status);
			sprint(w-&gt;msg, &#34;signal: %s&#34;, _p9sigstr(WTERMSIG(status), tmp));
			if(cd)
				strcat(w-&gt;msg, &#34; (core dumped)&#34;);
			return w;
		}
	}
}

Waitmsg*
p9wait(void)
{
	return _wait(-1, 0);
}

Waitmsg*
p9waitfor(int pid)
{
	return _wait(pid, 0);
}

Waitmsg*
p9waitnohang(void)
{
	return _wait(-1, WNOHANG);
}

int
p9waitpid(void)
{
	int status;
	return wait(&amp;status);
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

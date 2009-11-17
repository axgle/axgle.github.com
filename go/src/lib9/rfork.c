<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/rfork.c</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/lib9/rfork.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
/*
Plan 9 from User Space src/lib9/rfork.c
http://code.swtch.com/plan9port/src/tip/src/lib9/rfork.c

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
#include &lt;sys/wait.h&gt;
#include &lt;signal.h&gt;
#include &lt;libc.h&gt;
#undef rfork

static void
nop(int x)
{
	USED(x);
}

int
p9rfork(int flags)
{
	int pid, status;
	int p[2];
	int n;
	char buf[128], *q;
	extern char **environ;

	if((flags&amp;(RFPROC|RFFDG|RFMEM)) == (RFPROC|RFFDG)){
		/* check other flags before we commit */
		flags &amp;= ~(RFPROC|RFFDG|RFENVG);
		n = (flags &amp; ~(RFNOTEG|RFNAMEG|RFNOWAIT|RFCENVG));
		if(n){
			werrstr(&#34;unknown flags %08ux in rfork&#34;, n);
			return -1;
		}
		if(flags&amp;RFNOWAIT){
			/*
			 * BUG - should put the signal handler back after we
			 * finish, but I just don&#39;t care.  If a program calls with
			 * NOWAIT once, they&#39;re not likely to want child notes
			 * after that.
			 */
			signal(SIGCHLD, nop);
			if(pipe(p) &lt; 0)
				return -1;
		}
		pid = fork();
		if(pid == -1)
			return -1;
		if(flags&amp;RFNOWAIT){
			flags &amp;= ~RFNOWAIT;
			if(pid){
				/*
				 * Parent - wait for child to fork wait-free child.
				 * Then read pid from pipe.  Assume pipe buffer can absorb the write.
				 */
				close(p[1]);
				status = 0;
				if(wait4(pid, &amp;status, 0, 0) &lt; 0){
					werrstr(&#34;pipe dance - wait4 - %r&#34;);
					close(p[0]);
					return -1;
				}
				n = readn(p[0], buf, sizeof buf-1);
				close(p[0]);
				if(!WIFEXITED(status) || WEXITSTATUS(status)!=0 || n &lt;= 0){
					if(!WIFEXITED(status))
						werrstr(&#34;pipe dance - !exited 0x%ux&#34;, status);
					else if(WEXITSTATUS(status) != 0)
						werrstr(&#34;pipe dance - non-zero status 0x%ux&#34;, status);
					else if(n &lt; 0)
						werrstr(&#34;pipe dance - pipe read error - %r&#34;);
					else if(n == 0)
						werrstr(&#34;pipe dance - pipe read eof&#34;);
					else
						werrstr(&#34;pipe dance - unknown failure&#34;);
					return -1;
				}
				buf[n] = 0;
				if(buf[0] == &#39;x&#39;){
					werrstr(&#34;%s&#34;, buf+2);
					return -1;
				}
				pid = strtol(buf, &amp;q, 0);
			}else{
				/*
				 * Child - fork a new child whose wait message can&#39;t
				 * get back to the parent because we&#39;re going to exit!
				 */
				signal(SIGCHLD, SIG_IGN);
				close(p[0]);
				pid = fork();
				if(pid){
					/* Child parent - send status over pipe and exit. */
					if(pid &gt; 0)
						fprint(p[1], &#34;%d&#34;, pid);
					else
						fprint(p[1], &#34;x %r&#34;);
					close(p[1]);
					_exit(0);
				}else{
					/* Child child - close pipe. */
					close(p[1]);
				}
			}
		}
		if(pid != 0)
			return pid;
		if(flags&amp;RFCENVG)
			if(environ)
				*environ = nil;
	}
	if(flags&amp;RFPROC){
		werrstr(&#34;cannot use rfork for shared memory -- use libthread&#34;);
		return -1;
	}
	if(flags&amp;RFNAMEG){
		/* XXX set $NAMESPACE to a new directory */
		flags &amp;= ~RFNAMEG;
	}
	if(flags&amp;RFNOTEG){
		setpgid(0, getpid());
		flags &amp;= ~RFNOTEG;
	}
	if(flags&amp;RFNOWAIT){
		werrstr(&#34;cannot use RFNOWAIT without RFPROC&#34;);
		return -1;
	}
	if(flags){
		werrstr(&#34;unknown flags %08ux in rfork&#34;, flags);
		return -1;
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

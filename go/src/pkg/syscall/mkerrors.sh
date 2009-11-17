<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/syscall/mkerrors.sh</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/syscall/mkerrors.sh</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
#!/bin/bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Generate Go code listing errors and other #defined constant
# values (ENAMETOOLONG etc.), by asking the preprocessor
# about the definitions.

case &#34;$GOARCH&#34; in
arm)
	GCC=arm-gcc
	;;
*)
	GCC=gcc
	;;
esac

uname=$(uname)

includes_Linux=&#39;
#define _LARGEFILE_SOURCE
#define _LARGEFILE64_SOURCE
#define _FILE_OFFSET_BITS 64
#define _GNU_SOURCE

#include &lt;sys/types.h&gt;
#include &lt;sys/epoll.h&gt;
#include &lt;linux/ptrace.h&gt;
#include &lt;linux/wait.h&gt;
&#39;

includes_Darwin=&#39;
#define __DARWIN_UNIX03 0
#define KERNEL
#define _DARWIN_USE_64_BIT_INODE
#include &lt;sys/wait.h&gt;
#include &lt;sys/event.h&gt;
&#39;

includes=&#39;
#include &lt;sys/types.h&gt;
#include &lt;fcntl.h&gt;
#include &lt;dirent.h&gt;
#include &lt;sys/socket.h&gt;
#include &lt;netinet/in.h&gt;
#include &lt;netinet/tcp.h&gt;
#include &lt;errno.h&gt;
#include &lt;sys/signal.h&gt;
#include &lt;signal.h&gt;
&#39;

# Write godefs input.
(
	indirect=&#34;includes_$(uname)&#34;
	echo &#34;${!indirect} $includes&#34;
	echo
	echo &#39;enum {&#39;

	# The gcc command line prints all the #defines
	# it encounters while processing the input
	echo &#34;${!indirect} $includes&#34; | $GCC -x c - -E -dM |
	awk &#39;
		$1 != &#34;#define&#34; || $2 ~ /\(/ {next}
		
		$2 ~ /^(SIGEV_|SIGSTKSZ|SIGRT(MIN|MAX))/ {next}

		$2 ~ /^E[A-Z0-9_]+$/ ||
		$2 ~ /^SIG[^_]/ ||
		$2 ~ /^(AF|SOCK|SO|SOL|IPPROTO|IP|TCP|EVFILT|EV)_/ ||
		$2 == &#34;SOMAXCONN&#34; ||
		$2 == &#34;NAME_MAX&#34; ||
		$2 ~ /^(O|F|FD|NAME|S|PTRACE)_/ ||
		$2 ~ /^W[A-Z0-9]+$/ {printf(&#34;\t$%s = %s,\n&#34;, $2, $2)}
		
		$2 ~ /^__W[A-Z0-9]+$/ {printf(&#34;\t$%s = %s,\n&#34;, substr($2,3), $2)}
		
		{next}
	&#39; | sort

	echo &#39;};&#39;
) &gt;_const.c

# Pull out just the error names for later.
errors=$(
	echo &#39;#include &lt;errno.h&gt;&#39; | $GCC -x c - -E -dM |
	awk &#39;$1==&#34;#define&#34; &amp;&amp; $2 ~ /^E[A-Z0-9_]+$/ { print $2 }&#39;
)

echo &#39;// mkerrors.sh&#39; &#34;$@&#34;
echo &#39;// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT&#39;
echo
godefs -gsyscall &#34;$@&#34; _const.c

# Run C program to print error strings.
(
	/bin/echo &#34;
#include &lt;stdio.h&gt;
#include &lt;errno.h&gt;
#include &lt;ctype.h&gt;
#include &lt;string.h&gt;

#define nelem(x) (sizeof(x)/sizeof((x)[0]))

enum { A = &#39;A&#39;, Z = &#39;Z&#39;, a = &#39;a&#39;, z = &#39;z&#39; }; // avoid need for single quotes below

int errors[] = {
&#34;
	for i in $errors
	do
		/bin/echo &#39;	&#39;$i,
	done

	# Use /bin/echo to avoid builtin echo,
	# which interprets \n itself
	/bin/echo &#39;
};

int
main(void)
{
	int i, j, e;
	char buf[1024];

	printf(&#34;\n\n// Error table\n&#34;);
	printf(&#34;var errors = [...]string {\n&#34;);
	for(i=0; i&lt;nelem(errors); i++) {
		e = errors[i];
		for(j=0; j&lt;i; j++)
			if(errors[j] == e)	// duplicate value
				goto next;
		strcpy(buf, strerror(e));
		// lowercase first letter: Bad -&gt; bad, but STREAM -&gt; STREAM.
		if(A &lt;= buf[0] &amp;&amp; buf[0] &lt;= Z &amp;&amp; a &lt;= buf[1] &amp;&amp; buf[1] &lt;= z)
			buf[0] += a - A;
		printf(&#34;\t%d: \&#34;%s\&#34;,\n&#34;, e, buf);
	next:;
	}
	printf(&#34;}\n\n&#34;);
	return 0;
}

&#39;
) &gt;_errors.c

gcc -o _errors _errors.c &amp;&amp; ./_errors &amp;&amp; rm -f _errors.c _errors _const.c
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

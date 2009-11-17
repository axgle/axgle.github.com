<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/make.bash</title>

  <link rel="stylesheet" type="text/css" href="../doc/style.css">
  <script type="text/javascript" src="../doc/godocs.js"></script>

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
        <a href="../index.html"><img src="../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../doc/go_faq.html">FAQ</a></li>
    <li><a href="../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../doc/install.html">Install Go</a></li>
    <li><a href="../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../cmd/index.html">Command documentation</a></li>
    <li><a href="../pkg/index.html">Package documentation</a></li>
    <li><a href="index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/make.bash</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
#!/bin/bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e
GOBIN=&#34;${GOBIN:-$HOME/bin}&#34;
export MAKEFLAGS=-j4

if ! test -f $GOROOT/include/u.h
then
	echo &#39;$GOROOT is not set correctly or not exported&#39; 1&gt;&amp;2
	exit 1
fi

if ! test -d $GOBIN
then
	echo &#39;$GOBIN is not a directory or does not exist&#39; 1&gt;&amp;2
	echo &#39;create it or set $GOBIN differently&#39; 1&gt;&amp;2
	exit 1
fi

case &#34;$GOARCH&#34; in
amd64 | 386 | arm)
	;;
*)
	echo &#39;$GOARCH is set to &lt;&#39;$GOARCH&#39;&gt;, must be amd64, 386, or arm&#39; 1&gt;&amp;2
	exit 1
esac

case &#34;$GOOS&#34; in
darwin | linux | nacl)
	;;
*)
	echo &#39;$GOOS is set to &lt;&#39;$GOOS&#39;&gt;, must be darwin, linux, or nacl&#39; 1&gt;&amp;2
	exit 1
esac

rm -f $GOBIN/quietgcc
CC=${CC:-gcc}
sed -e &#34;s|@CC@|$CC|&#34; &lt; quietgcc.bash &gt; $GOBIN/quietgcc
chmod +x $GOBIN/quietgcc

if ! (cd lib9 &amp;&amp; which quietgcc) &gt;/dev/null 2&gt;&amp;1; then
	echo &#34;installed quietgcc as $GOBIN/quietgcc but &#39;which quietgcc&#39; fails&#34; 1&gt;&amp;2
	echo &#34;double-check that $GOBIN is in your &#34;&#39;$PATH&#39; 1&gt;&amp;2
	exit 1
fi

if [ -d /selinux -a -f /selinux/booleans/allow_execstack ] ; then
	if ! cat /selinux/booleans/allow_execstack | grep -c &#39;^1 1$&#39; &gt;&gt; /dev/null ; then
		echo &#34;WARNING: the default SELinux policy on, at least, Fedora 12 breaks &#34;
		echo &#34;Go. You can enable the features that Go needs via the following &#34;
		echo &#34;command (as root):&#34;
		echo &#34;  # setsebool -P allow_execstack 1&#34;
		echo
		echo &#34;Note that this affects your system globally! &#34;
		echo
		echo &#34;The build will continue in five seconds in case we &#34;
		echo &#34;misdiagnosed the issue...&#34;

		sleep 5
	fi
fi

bash clean.bash

for i in lib9 libbio libmach cmd pkg libcgo cmd/cgo cmd/ebnflint cmd/godoc cmd/gofmt cmd/goyacc cmd/hgpatch
do
	case &#34;$i-$GOOS&#34; in
	libcgo-nacl)
		;;
	*)
		# The ( ) here are to preserve the current directory
		# for the next round despite the cd $i below.
		# set -e does not apply to ( ) so we must explicitly
		# test the exit status.
		(
			echo; echo; echo %%%% making $i %%%%; echo
			cd $i
			case $i in
			cmd)
				bash make.bash
				;;
			*)
				make install
			esac
		)  || exit 1
	esac
done

case &#34;`uname`&#34; in
Darwin)
	echo;
	echo %%% run sudo.bash to install debuggers
	echo
esac
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

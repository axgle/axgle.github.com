<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gofmt/test.sh</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gofmt/test.sh</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
#!/bin/bash
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

. $GOROOT/src/Make.$GOARCH
if [ -z &#34;$O&#34; ]; then
	echo &#39;missing $O - maybe no Make.$GOARCH?&#39; 1&gt;&amp;2
	exit 1
fi

CMD=&#34;./gofmt&#34;
TMP1=test_tmp1.go
TMP2=test_tmp2.go
TMP3=test_tmp3.go
COUNT=0

count() {
	#echo $1
	let COUNT=$COUNT+1
	let M=$COUNT%10
	if [ $M == 0 ]; then
		echo -n &#34;.&#34;
	fi
}


# apply to one file
apply1() {
	#echo $1 $2
	case `basename $F` in
	# except for elf.go (which is not yet idempotent due to a few
	# tricky-to-format comments) the following files are skipped
	# because they are test cases for syntax errors and thus won&#39;t
	# parse in the first place:
	elf.go | \
	func3.go | const2.go | \
	bug014.go | bug050.go |  bug068.go |  bug083.go | bug088.go | \
	bug106.go | bug121.go | bug125.go | bug133.go | bug160.go | \
	bug163.go | bug166.go | bug169.go ) ;;
	* ) $1 $2; count $F;;
	esac
}


# apply to local files
applydot() {
	for F in `find . -name &#34;*.go&#34; | grep -v &#34;._&#34;`; do
		apply1 $1 $F
	done
}


# apply to all .go files we can find
apply() {
	for F in `find $GOROOT -name &#34;*.go&#34; | grep -v &#34;._&#34;`; do
		apply1 $1 $F
	done
}


cleanup() {
	rm -f $TMP1 $TMP2 $TMP3
}


silent() {
	cleanup
	$CMD $1 &gt; /dev/null 2&gt; $TMP1
	if [ $? != 0 ]; then
		cat $TMP1
		echo &#34;Error (silent mode test): test.sh $1&#34;
		exit 1
	fi
}


idempotent() {
	cleanup
	$CMD $1 &gt; $TMP1
	if [ $? != 0 ]; then
		echo &#34;Error (step 1 of idempotency test): test.sh $1&#34;
		exit 1
	fi

	$CMD $TMP1 &gt; $TMP2
	if [ $? != 0 ]; then
		echo &#34;Error (step 2 of idempotency test): test.sh $1&#34;
		exit 1
	fi

	$CMD $TMP2 &gt; $TMP3
	if [ $? != 0 ]; then
		echo &#34;Error (step 3 of idempotency test): test.sh $1&#34;
		exit 1
	fi

	cmp -s $TMP2 $TMP3
	if [ $? != 0 ]; then
		diff $TMP2 $TMP3
		echo &#34;Error (step 4 of idempotency test): test.sh $1&#34;
		exit 1
	fi
}


valid() {
	cleanup
	$CMD $1 &gt; $TMP1
	if [ $? != 0 ]; then
		echo &#34;Error (step 1 of validity test): test.sh $1&#34;
		exit 1
	fi

	$GC -o /dev/null $TMP1
	if [ $? != 0 ]; then
		echo &#34;Error (step 2 of validity test): test.sh $1&#34;
		exit 1
	fi
}


runtest() {
	#echo &#34;Testing silent mode&#34;
	cleanup
	$1 silent $2

	#echo &#34;Testing idempotency&#34;
	cleanup
	$1 idempotent $2
}


runtests() {
	if [ $# == 0 ]; then
		runtest apply
		# verify the pretty-printed files can be compiled with $GC again
		# do it in local directory only because of the prerequisites required
		#echo &#34;Testing validity&#34;
		cleanup
		applydot valid
	else
		for F in $*; do
			runtest apply1 $F
		done
	fi
}


# run over all .go files
runtests $*
cleanup

# done
echo
echo &#34;PASSED ($COUNT tests)&#34;
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

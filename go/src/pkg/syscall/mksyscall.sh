<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/syscall/mksyscall.sh</title>

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
  <h1 id="generatedHeader">Text file src/pkg/syscall/mksyscall.sh</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
#!/usr/bin/perl
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# This program reads a file containing function prototypes
# (like syscall_darwin.go) and generates system call bodies.
# The prototypes are marked by lines beginning with &#34;//sys&#34;
# and read like func declarations if //sys is replaced by func, but:
#	* The parameter lists must give a name for each argument.
#	  This includes return parameters.
#	* The parameter lists must give a type for each argument:
#	  the (x, y, z int) shorthand is not allowed.
#	* If the return parameter is an error number, it must be named errno.

$cmdline = &#34;mksyscall.sh &#34; . join(&#39; &#39;, @ARGV);
$errors = 0;
$_32bit = &#34;&#34;;

if($ARGV[0] eq &#34;-b32&#34;) {
	$_32bit = &#34;big-endian&#34;;
	shift;
} elsif($ARGV[0] eq &#34;-l32&#34;) {
	$_32bit = &#34;little-endian&#34;;
	shift;
}

if($ARGV[0] =~ /^-/) {
	print STDERR &#34;usage: mksyscall.sh [-b32 | -l32] [file ...]\n&#34;;
	exit 1;
}

sub parseparamlist($) {
	my ($list) = @_;
	$list =~ s/^\s*//;
	$list =~ s/\s*$//;
	if($list eq &#34;&#34;) {
		return ();
	}
	return split(/\s*,\s*/, $list);
}

sub parseparam($) {
	my ($p) = @_;
	if($p !~ /^(\S*) (\S*)$/) {
		print STDERR &#34;$ARGV:$.: malformed parameter: $p\n&#34;;
		$errors = 1;
		return (&#34;xx&#34;, &#34;int&#34;);
	}
	return ($1, $2);
}

$text = &#34;&#34;;
while(&lt;&gt;) {
	chomp;
	s/\s+/ /g;
	s/^\s+//;
	s/\s+$//;
	next if !/^\/\/sys /;

	# Line must be of the form
	#	func Open(path string, mode int, perm int) (fd int, errno int)
	# Split into name, in params, out params.
	if(!/^\/\/sys (\w+)\(([^()]*)\)\s*(?:\(([^()]+)\))?\s*(?:=\s*(SYS_[A-Z0-9_]+))?$/) {
		print STDERR &#34;$ARGV:$.: malformed //sys declaration\n&#34;;
		$errors = 1;
		next;
	}
	my ($func, $in, $out, $sysname) = ($1, $2, $3, $4);

	# Split argument lists on comma.
	my @in = parseparamlist($in);
	my @out = parseparamlist($out);

	# Go function header.
	$text .= sprintf &#34;func %s(%s) (%s) {\n&#34;, $func, join(&#39;, &#39;, @in), join(&#39;, &#39;, @out);

	# Prepare arguments to Syscall.
	my @args = ();
	my $n = 0;
	foreach my $p (@in) {
		my ($name, $type) = parseparam($p);
		if($type =~ /^\*/) {
			push @args, &#34;uintptr(unsafe.Pointer($name))&#34;;
		} elsif($type eq &#34;string&#34;) {
			push @args, &#34;uintptr(unsafe.Pointer(StringBytePtr($name)))&#34;;
		} elsif($type =~ /^\[\](.*)/) {
			# Convert slice into pointer, length.
			# Have to be careful not to take address of &amp;a[0] if len == 0:
			# pass nil in that case.
			$text .= &#34;\tvar _p$n *$1;\n&#34;;
			$text .= &#34;\tif len($name) &gt; 0 { _p$n = \&amp;${name}[0]; }\n&#34;;
			push @args, &#34;uintptr(unsafe.Pointer(_p$n))&#34;, &#34;uintptr(len($name))&#34;;
			$n++;
		} elsif($type eq &#34;int64&#34; &amp;&amp; $_32bit ne &#34;&#34;) {
			if($_32bit eq &#34;big-endian&#34;) {
				push @args, &#34;uintptr($name &gt;&gt; 32)&#34;, &#34;uintptr($name)&#34;;
			} else {
				push @args, &#34;uintptr($name)&#34;, &#34;uintptr($name &gt;&gt; 32)&#34;;
			}
		} else {
			push @args, &#34;uintptr($name)&#34;;
		}
	}

	# Determine which form to use; pad args with zeros.
	my $asm = &#34;Syscall&#34;;
	if(@args &lt;= 3) {
		while(@args &lt; 3) {
			push @args, &#34;0&#34;;
		}
	} elsif(@args &lt;= 6) {
		$asm = &#34;Syscall6&#34;;
		while(@args &lt; 6) {
			push @args, &#34;0&#34;;
		}
	} else {
		print STDERR &#34;$ARGV:$.: too many arguments to system call\n&#34;;
	}

	# System call number.
	if($sysname eq &#34;&#34;) {
		$sysname = &#34;SYS_$func&#34;;
		$sysname =~ s/([a-z])([A-Z])/${1}_$2/g;	# turn FooBar into Foo_Bar
		$sysname =~ y/a-z/A-Z/;
	}

	# Actual call.
	my $args = join(&#39;, &#39;, @args);
	my $call = &#34;$asm($sysname, $args)&#34;;

	# Assign return values.
	my $body = &#34;&#34;;
	my @ret = (&#34;_&#34;, &#34;_&#34;, &#34;_&#34;);
	for(my $i=0; $i&lt;@out; $i++) {
		my $p = $out[$i];
		my ($name, $type) = parseparam($p);
		my $reg = &#34;&#34;;
		if($name eq &#34;errno&#34;) {
			$reg = &#34;e1&#34;;
			$ret[2] = $reg;
		} else {
			$reg = sprintf(&#34;r%d&#34;, $i);
			$ret[$i] = $reg;
		}
		if($type eq &#34;bool&#34;) {
			$reg = &#34;$reg != 0&#34;;
		}
		if($type eq &#34;int64&#34; &amp;&amp; $_32bit ne &#34;&#34;) {
			# 64-bit number in r1:r0 or r0:r1.
			if($i+2 &gt; @out) {
				print STDERR &#34;$ARGV:$.: not enough registers for int64 return\n&#34;;
			}
			if($_32bit eq &#34;big-endian&#34;) {
				$reg = sprintf(&#34;int64(r%d)&lt;&lt;32 | int64(r%d)&#34;, $i, $i+1);
			} else {
				$reg = sprintf(&#34;int64(r%d)&lt;&lt;32 | int64(r%d)&#34;, $i+1, $i);
			}
			$ret[$i] = sprintf(&#34;r%d&#34;, $i);
			$ret[$i+1] = sprintf(&#34;r%d&#34;, $i+1);
			$i++;		# loop will do another $i++
		}
		$body .= &#34;\t$name = $type($reg);\n&#34;;
	}
	if ($ret[0] eq &#34;_&#34; &amp;&amp; $ret[1] eq &#34;_&#34; &amp;&amp; $ret[2] eq &#34;_&#34;) {
		$text .= &#34;\t$call;\n&#34;;
	} else {
		$text .= &#34;\t$ret[0], $ret[1], $ret[2] := $call;\n&#34;;
	}
	$text .= $body;

	$text .= &#34;\treturn;\n&#34;;
	$text .= &#34;}\n\n&#34;;
}

if($errors) {
	exit 1;
}

print &lt;&lt;EOF;
// $cmdline
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

import &#34;unsafe&#34;

$text

EOF
exit 0;
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

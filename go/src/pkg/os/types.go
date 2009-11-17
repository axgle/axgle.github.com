<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/types.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/os/types.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package os

<a id="L7"></a>import &#34;syscall&#34;

<a id="L9"></a><span class="comment">// An operating-system independent representation of Unix data structures.</span>
<a id="L10"></a><span class="comment">// OS-specific routines in this directory convert the OS-local versions to these.</span>

<a id="L12"></a><span class="comment">// Getpagesize returns the underlying system&#39;s memory page size.</span>
<a id="L13"></a>func Getpagesize() int { return syscall.Getpagesize() }

<a id="L15"></a><span class="comment">// A Dir describes a file and is returned by Stat, Fstat, and Lstat</span>
<a id="L16"></a>type Dir struct {
    <a id="L17"></a>Dev             uint64; <span class="comment">// device number of file system holding file.</span>
    <a id="L18"></a>Ino             uint64; <span class="comment">// inode number.</span>
    <a id="L19"></a>Nlink           uint64; <span class="comment">// number of hard links.</span>
    <a id="L20"></a>Mode            uint32; <span class="comment">// permission and mode bits.</span>
    <a id="L21"></a>Uid             uint32; <span class="comment">// user id of owner.</span>
    <a id="L22"></a>Gid             uint32; <span class="comment">// group id of owner.</span>
    <a id="L23"></a>Rdev            uint64; <span class="comment">// device type for special file.</span>
    <a id="L24"></a>Size            uint64; <span class="comment">// length in bytes.</span>
    <a id="L25"></a>Blksize         uint64; <span class="comment">// size of blocks, in bytes.</span>
    <a id="L26"></a>Blocks          uint64; <span class="comment">// number of blocks allocated for file.</span>
    <a id="L27"></a>Atime_ns        uint64; <span class="comment">// access time; nanoseconds since epoch.</span>
    <a id="L28"></a>Mtime_ns        uint64; <span class="comment">// modified time; nanoseconds since epoch.</span>
    <a id="L29"></a>Ctime_ns        uint64; <span class="comment">// status change time; nanoseconds since epoch.</span>
    <a id="L30"></a>Name            string; <span class="comment">// name of file as presented to Open.</span>
    <a id="L31"></a>FollowedSymlink bool;   <span class="comment">// followed a symlink to get this information</span>
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// IsFifo reports whether the Dir describes a FIFO file.</span>
<a id="L35"></a>func (dir *Dir) IsFifo() bool { return (dir.Mode &amp; syscall.S_IFMT) == syscall.S_IFIFO }

<a id="L37"></a><span class="comment">// IsChar reports whether the Dir describes a character special file.</span>
<a id="L38"></a>func (dir *Dir) IsChar() bool { return (dir.Mode &amp; syscall.S_IFMT) == syscall.S_IFCHR }

<a id="L40"></a><span class="comment">// IsDirectory reports whether the Dir describes a directory.</span>
<a id="L41"></a>func (dir *Dir) IsDirectory() bool { return (dir.Mode &amp; syscall.S_IFMT) == syscall.S_IFDIR }

<a id="L43"></a><span class="comment">// IsBlock reports whether the Dir describes a block special file.</span>
<a id="L44"></a>func (dir *Dir) IsBlock() bool { return (dir.Mode &amp; syscall.S_IFMT) == syscall.S_IFBLK }

<a id="L46"></a><span class="comment">// IsRegular reports whether the Dir describes a regular file.</span>
<a id="L47"></a>func (dir *Dir) IsRegular() bool { return (dir.Mode &amp; syscall.S_IFMT) == syscall.S_IFREG }

<a id="L49"></a><span class="comment">// IsSymlink reports whether the Dir describes a symbolic link.</span>
<a id="L50"></a>func (dir *Dir) IsSymlink() bool { return (dir.Mode &amp; syscall.S_IFMT) == syscall.S_IFLNK }

<a id="L52"></a><span class="comment">// IsSocket reports whether the Dir describes a socket.</span>
<a id="L53"></a>func (dir *Dir) IsSocket() bool { return (dir.Mode &amp; syscall.S_IFMT) == syscall.S_IFSOCK }

<a id="L55"></a><span class="comment">// Permission returns the file permission bits.</span>
<a id="L56"></a>func (dir *Dir) Permission() int { return int(dir.Mode &amp; 0777) }
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

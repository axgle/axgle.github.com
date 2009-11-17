<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/error.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/os/error.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package os

<a id="L7"></a>import syscall &#34;syscall&#34;

<a id="L9"></a><span class="comment">// An Error can represent any printable error condition.</span>
<a id="L10"></a>type Error interface {
    <a id="L11"></a>String() string;
<a id="L12"></a>}

<a id="L14"></a><span class="comment">// A helper type that can be embedded or wrapped to simplify satisfying</span>
<a id="L15"></a><span class="comment">// Error.</span>
<a id="L16"></a>type ErrorString string

<a id="L18"></a>func (e ErrorString) String() string { return string(e) }

<a id="L20"></a><span class="comment">// Note: If the name of the function NewError changes,</span>
<a id="L21"></a><span class="comment">// pkg/go/doc/doc.go should be adjusted since it hardwires</span>
<a id="L22"></a><span class="comment">// this name in a heuristic.</span>

<a id="L24"></a><span class="comment">// NewError converts s to an ErrorString, which satisfies the Error interface.</span>
<a id="L25"></a>func NewError(s string) Error { return ErrorString(s) }

<a id="L27"></a><span class="comment">// Errno is the Unix error number.  Names such as EINVAL are simple</span>
<a id="L28"></a><span class="comment">// wrappers to convert the error number into an Error.</span>
<a id="L29"></a>type Errno int64

<a id="L31"></a>func (e Errno) String() string { return syscall.Errstr(int(e)) }

<a id="L33"></a><span class="comment">// Commonly known Unix errors.</span>
<a id="L34"></a>var (
    <a id="L35"></a>EPERM        Error = Errno(syscall.EPERM);
    <a id="L36"></a>ENOENT       Error = Errno(syscall.ENOENT);
    <a id="L37"></a>ESRCH        Error = Errno(syscall.ESRCH);
    <a id="L38"></a>EINTR        Error = Errno(syscall.EINTR);
    <a id="L39"></a>EIO          Error = Errno(syscall.EIO);
    <a id="L40"></a>ENXIO        Error = Errno(syscall.ENXIO);
    <a id="L41"></a>E2BIG        Error = Errno(syscall.E2BIG);
    <a id="L42"></a>ENOEXEC      Error = Errno(syscall.ENOEXEC);
    <a id="L43"></a>EBADF        Error = Errno(syscall.EBADF);
    <a id="L44"></a>ECHILD       Error = Errno(syscall.ECHILD);
    <a id="L45"></a>EDEADLK      Error = Errno(syscall.EDEADLK);
    <a id="L46"></a>ENOMEM       Error = Errno(syscall.ENOMEM);
    <a id="L47"></a>EACCES       Error = Errno(syscall.EACCES);
    <a id="L48"></a>EFAULT       Error = Errno(syscall.EFAULT);
    <a id="L49"></a>EBUSY        Error = Errno(syscall.EBUSY);
    <a id="L50"></a>EEXIST       Error = Errno(syscall.EEXIST);
    <a id="L51"></a>EXDEV        Error = Errno(syscall.EXDEV);
    <a id="L52"></a>ENODEV       Error = Errno(syscall.ENODEV);
    <a id="L53"></a>ENOTDIR      Error = Errno(syscall.ENOTDIR);
    <a id="L54"></a>EISDIR       Error = Errno(syscall.EISDIR);
    <a id="L55"></a>EINVAL       Error = Errno(syscall.EINVAL);
    <a id="L56"></a>ENFILE       Error = Errno(syscall.ENFILE);
    <a id="L57"></a>EMFILE       Error = Errno(syscall.EMFILE);
    <a id="L58"></a>ENOTTY       Error = Errno(syscall.ENOTTY);
    <a id="L59"></a>EFBIG        Error = Errno(syscall.EFBIG);
    <a id="L60"></a>ENOSPC       Error = Errno(syscall.ENOSPC);
    <a id="L61"></a>ESPIPE       Error = Errno(syscall.ESPIPE);
    <a id="L62"></a>EROFS        Error = Errno(syscall.EROFS);
    <a id="L63"></a>EMLINK       Error = Errno(syscall.EMLINK);
    <a id="L64"></a>EPIPE        Error = Errno(syscall.EPIPE);
    <a id="L65"></a>EAGAIN       Error = Errno(syscall.EAGAIN);
    <a id="L66"></a>EDOM         Error = Errno(syscall.EDOM);
    <a id="L67"></a>ERANGE       Error = Errno(syscall.ERANGE);
    <a id="L68"></a>EADDRINUSE   Error = Errno(syscall.EADDRINUSE);
    <a id="L69"></a>ECONNREFUSED Error = Errno(syscall.ECONNREFUSED);
    <a id="L70"></a>ENAMETOOLONG Error = Errno(syscall.ENAMETOOLONG);
    <a id="L71"></a>EAFNOSUPPORT Error = Errno(syscall.EAFNOSUPPORT);
<a id="L72"></a>)

<a id="L74"></a><span class="comment">// PathError records an error and the operation and file path that caused it.</span>
<a id="L75"></a>type PathError struct {
    <a id="L76"></a>Op    string;
    <a id="L77"></a>Path  string;
    <a id="L78"></a>Error Error;
<a id="L79"></a>}

<a id="L81"></a>func (e *PathError) String() string { return e.Op + &#34; &#34; + e.Path + &#34;: &#34; + e.Error.String() }

<a id="L83"></a><span class="comment">// SyscallError records an error from a specific system call.</span>
<a id="L84"></a>type SyscallError struct {
    <a id="L85"></a>Syscall string;
    <a id="L86"></a>Errno   Errno;
<a id="L87"></a>}

<a id="L89"></a>func (e *SyscallError) String() string { return e.Syscall + &#34;: &#34; + e.Errno.String() }

<a id="L91"></a><span class="comment">// Note: If the name of the function NewSyscallError changes,</span>
<a id="L92"></a><span class="comment">// pkg/go/doc/doc.go should be adjusted since it hardwires</span>
<a id="L93"></a><span class="comment">// this name in a heuristic.</span>

<a id="L95"></a><span class="comment">// NewSyscallError returns, as an Error, a new SyscallError</span>
<a id="L96"></a><span class="comment">// with the given system call name and error number.</span>
<a id="L97"></a><span class="comment">// As a convenience, if errno is 0, NewSyscallError returns nil.</span>
<a id="L98"></a>func NewSyscallError(syscall string, errno int) Error {
    <a id="L99"></a>if errno == 0 {
        <a id="L100"></a>return nil
    <a id="L101"></a>}
    <a id="L102"></a>return &amp;SyscallError{syscall, Errno(errno)};
<a id="L103"></a>}
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

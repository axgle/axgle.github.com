<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/file.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/os/file.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The os package provides a platform-independent interface to operating</span>
<a id="L6"></a><span class="comment">// system functionality.  The design is Unix-like.</span>
<a id="L7"></a>package os

<a id="L9"></a>import (
    <a id="L10"></a>&#34;syscall&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// Auxiliary information if the File describes a directory</span>
<a id="L14"></a>type dirInfo struct {
    <a id="L15"></a>buf  []byte; <span class="comment">// buffer for directory I/O</span>
    <a id="L16"></a>nbuf int;    <span class="comment">// length of buf; return value from Getdirentries</span>
    <a id="L17"></a>bufp int;    <span class="comment">// location of next record in buf.</span>
<a id="L18"></a>}

<a id="L20"></a><span class="comment">// File represents an open file descriptor.</span>
<a id="L21"></a>type File struct {
    <a id="L22"></a>fd      int;
    <a id="L23"></a>name    string;
    <a id="L24"></a>dirinfo *dirInfo; <span class="comment">// nil unless directory being read</span>
    <a id="L25"></a>nepipe  int;      <span class="comment">// number of consecutive EPIPE in Write</span>
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// Fd returns the integer Unix file descriptor referencing the open file.</span>
<a id="L29"></a>func (file *File) Fd() int { return file.fd }

<a id="L31"></a><span class="comment">// Name returns the name of the file as presented to Open.</span>
<a id="L32"></a>func (file *File) Name() string { return file.name }

<a id="L34"></a><span class="comment">// NewFile returns a new File with the given file descriptor and name.</span>
<a id="L35"></a>func NewFile(fd int, name string) *File {
    <a id="L36"></a>if fd &lt; 0 {
        <a id="L37"></a>return nil
    <a id="L38"></a>}
    <a id="L39"></a>return &amp;File{fd, name, nil, 0};
<a id="L40"></a>}

<a id="L42"></a><span class="comment">// Stdin, Stdout, and Stderr are open Files pointing to the standard input,</span>
<a id="L43"></a><span class="comment">// standard output, and standard error file descriptors.</span>
<a id="L44"></a>var (
    <a id="L45"></a>Stdin  = NewFile(0, &#34;/dev/stdin&#34;);
    <a id="L46"></a>Stdout = NewFile(1, &#34;/dev/stdout&#34;);
    <a id="L47"></a>Stderr = NewFile(2, &#34;/dev/stderr&#34;);
<a id="L48"></a>)

<a id="L50"></a><span class="comment">// Flags to Open wrapping those of the underlying system. Not all flags</span>
<a id="L51"></a><span class="comment">// may be implemented on a given system.</span>
<a id="L52"></a>const (
    <a id="L53"></a>O_RDONLY   = syscall.O_RDONLY;   <span class="comment">// open the file read-only.</span>
    <a id="L54"></a>O_WRONLY   = syscall.O_WRONLY;   <span class="comment">// open the file write-only.</span>
    <a id="L55"></a>O_RDWR     = syscall.O_RDWR;     <span class="comment">// open the file read-write.</span>
    <a id="L56"></a>O_APPEND   = syscall.O_APPEND;   <span class="comment">// open the file append-only.</span>
    <a id="L57"></a>O_ASYNC    = syscall.O_ASYNC;    <span class="comment">// generate a signal when I/O is available.</span>
    <a id="L58"></a>O_CREAT    = syscall.O_CREAT;    <span class="comment">// create a new file if none exists.</span>
    <a id="L59"></a>O_EXCL     = syscall.O_EXCL;     <span class="comment">// used with O_CREAT, file must not exist</span>
    <a id="L60"></a>O_NOCTTY   = syscall.O_NOCTTY;   <span class="comment">// do not make file the controlling tty.</span>
    <a id="L61"></a>O_NONBLOCK = syscall.O_NONBLOCK; <span class="comment">// open in non-blocking mode.</span>
    <a id="L62"></a>O_NDELAY   = O_NONBLOCK;         <span class="comment">// synonym for O_NONBLOCK</span>
    <a id="L63"></a>O_SYNC     = syscall.O_SYNC;     <span class="comment">// open for synchronous I/O.</span>
    <a id="L64"></a>O_TRUNC    = syscall.O_TRUNC;    <span class="comment">// if possible, truncate file when opened.</span>
    <a id="L65"></a>O_CREATE   = O_CREAT;            <span class="comment">// create a new file if none exists.</span>
<a id="L66"></a>)

<a id="L68"></a><span class="comment">// Open opens the named file with specified flag (O_RDONLY etc.) and perm, (0666 etc.)</span>
<a id="L69"></a><span class="comment">// if applicable.  If successful, methods on the returned File can be used for I/O.</span>
<a id="L70"></a><span class="comment">// It returns the File and an Error, if any.</span>
<a id="L71"></a>func Open(name string, flag int, perm int) (file *File, err Error) {
    <a id="L72"></a>r, e := syscall.Open(name, flag|syscall.O_CLOEXEC, perm);
    <a id="L73"></a>if e != 0 {
        <a id="L74"></a>return nil, &amp;PathError{&#34;open&#34;, name, Errno(e)}
    <a id="L75"></a>}

    <a id="L77"></a><span class="comment">// There&#39;s a race here with fork/exec, which we are</span>
    <a id="L78"></a><span class="comment">// content to live with.  See ../syscall/exec.go</span>
    <a id="L79"></a>if syscall.O_CLOEXEC == 0 { <span class="comment">// O_CLOEXEC not supported</span>
        <a id="L80"></a>syscall.CloseOnExec(r)
    <a id="L81"></a>}

    <a id="L83"></a>return NewFile(r, name), nil;
<a id="L84"></a>}

<a id="L86"></a><span class="comment">// Close closes the File, rendering it unusable for I/O.</span>
<a id="L87"></a><span class="comment">// It returns an Error, if any.</span>
<a id="L88"></a>func (file *File) Close() Error {
    <a id="L89"></a>if file == nil {
        <a id="L90"></a>return EINVAL
    <a id="L91"></a>}
    <a id="L92"></a>var err Error;
    <a id="L93"></a>if e := syscall.Close(file.fd); e != 0 {
        <a id="L94"></a>err = &amp;PathError{&#34;close&#34;, file.name, Errno(e)}
    <a id="L95"></a>}
    <a id="L96"></a>file.fd = -1; <span class="comment">// so it can&#39;t be closed again</span>
    <a id="L97"></a>return err;
<a id="L98"></a>}

<a id="L100"></a>type eofError int

<a id="L102"></a>func (eofError) String() string { return &#34;EOF&#34; }

<a id="L104"></a><span class="comment">// EOF is the Error returned by Read when no more input is available.</span>
<a id="L105"></a><span class="comment">// Functions should return EOF only to signal a graceful end of input.</span>
<a id="L106"></a><span class="comment">// If the EOF occurs unexpectedly in a structured data stream,</span>
<a id="L107"></a><span class="comment">// the appropriate error is either io.ErrUnexpectedEOF or some other error</span>
<a id="L108"></a><span class="comment">// giving more detail.</span>
<a id="L109"></a>var EOF Error = eofError(0)

<a id="L111"></a><span class="comment">// Read reads up to len(b) bytes from the File.</span>
<a id="L112"></a><span class="comment">// It returns the number of bytes read and an Error, if any.</span>
<a id="L113"></a><span class="comment">// EOF is signaled by a zero count with err set to EOF.</span>
<a id="L114"></a>func (file *File) Read(b []byte) (n int, err Error) {
    <a id="L115"></a>if file == nil {
        <a id="L116"></a>return 0, EINVAL
    <a id="L117"></a>}
    <a id="L118"></a>n, e := syscall.Read(file.fd, b);
    <a id="L119"></a>if n &lt; 0 {
        <a id="L120"></a>n = 0
    <a id="L121"></a>}
    <a id="L122"></a>if n == 0 &amp;&amp; e == 0 {
        <a id="L123"></a>return 0, EOF
    <a id="L124"></a>}
    <a id="L125"></a>if e != 0 {
        <a id="L126"></a>err = &amp;PathError{&#34;read&#34;, file.name, Errno(e)}
    <a id="L127"></a>}
    <a id="L128"></a>return n, err;
<a id="L129"></a>}

<a id="L131"></a><span class="comment">// ReadAt reads len(b) bytes from the File starting at byte offset off.</span>
<a id="L132"></a><span class="comment">// It returns the number of bytes read and the Error, if any.</span>
<a id="L133"></a><span class="comment">// EOF is signaled by a zero count with err set to EOF.</span>
<a id="L134"></a><span class="comment">// ReadAt always returns a non-nil Error when n != len(b).</span>
<a id="L135"></a>func (file *File) ReadAt(b []byte, off int64) (n int, err Error) {
    <a id="L136"></a>if file == nil {
        <a id="L137"></a>return 0, EINVAL
    <a id="L138"></a>}
    <a id="L139"></a>for len(b) &gt; 0 {
        <a id="L140"></a>m, e := syscall.Pread(file.fd, b, off);
        <a id="L141"></a>n += m;
        <a id="L142"></a>if e != 0 {
            <a id="L143"></a>err = &amp;PathError{&#34;read&#34;, file.name, Errno(e)};
            <a id="L144"></a>break;
        <a id="L145"></a>}
        <a id="L146"></a>b = b[m:len(b)];
        <a id="L147"></a>off += int64(m);
    <a id="L148"></a>}
    <a id="L149"></a>return;
<a id="L150"></a>}

<a id="L152"></a><span class="comment">// Write writes len(b) bytes to the File.</span>
<a id="L153"></a><span class="comment">// It returns the number of bytes written and an Error, if any.</span>
<a id="L154"></a><span class="comment">// Write returns a non-nil Error when n != len(b).</span>
<a id="L155"></a>func (file *File) Write(b []byte) (n int, err Error) {
    <a id="L156"></a>if file == nil {
        <a id="L157"></a>return 0, EINVAL
    <a id="L158"></a>}
    <a id="L159"></a>n, e := syscall.Write(file.fd, b);
    <a id="L160"></a>if n &lt; 0 {
        <a id="L161"></a>n = 0
    <a id="L162"></a>}
    <a id="L163"></a>if e == syscall.EPIPE {
        <a id="L164"></a>file.nepipe++;
        <a id="L165"></a>if file.nepipe &gt;= 10 {
            <a id="L166"></a>Exit(syscall.EPIPE)
        <a id="L167"></a>}
    <a id="L168"></a>} else {
        <a id="L169"></a>file.nepipe = 0
    <a id="L170"></a>}
    <a id="L171"></a>if e != 0 {
        <a id="L172"></a>err = &amp;PathError{&#34;write&#34;, file.name, Errno(e)}
    <a id="L173"></a>}
    <a id="L174"></a>return n, err;
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// WriteAt writes len(b) bytes to the File starting at byte offset off.</span>
<a id="L178"></a><span class="comment">// It returns the number of bytes written and an Error, if any.</span>
<a id="L179"></a><span class="comment">// WriteAt returns a non-nil Error when n != len(b).</span>
<a id="L180"></a>func (file *File) WriteAt(b []byte, off int64) (n int, err Error) {
    <a id="L181"></a>if file == nil {
        <a id="L182"></a>return 0, EINVAL
    <a id="L183"></a>}
    <a id="L184"></a>for len(b) &gt; 0 {
        <a id="L185"></a>m, e := syscall.Pwrite(file.fd, b, off);
        <a id="L186"></a>n += m;
        <a id="L187"></a>if e != 0 {
            <a id="L188"></a>err = &amp;PathError{&#34;write&#34;, file.name, Errno(e)};
            <a id="L189"></a>break;
        <a id="L190"></a>}
        <a id="L191"></a>b = b[m:len(b)];
        <a id="L192"></a>off += int64(m);
    <a id="L193"></a>}
    <a id="L194"></a>return;
<a id="L195"></a>}

<a id="L197"></a><span class="comment">// Seek sets the offset for the next Read or Write on file to offset, interpreted</span>
<a id="L198"></a><span class="comment">// according to whence: 0 means relative to the origin of the file, 1 means</span>
<a id="L199"></a><span class="comment">// relative to the current offset, and 2 means relative to the end.</span>
<a id="L200"></a><span class="comment">// It returns the new offset and an Error, if any.</span>
<a id="L201"></a>func (file *File) Seek(offset int64, whence int) (ret int64, err Error) {
    <a id="L202"></a>r, e := syscall.Seek(file.fd, offset, whence);
    <a id="L203"></a>if e == 0 &amp;&amp; file.dirinfo != nil &amp;&amp; r != 0 {
        <a id="L204"></a>e = syscall.EISDIR
    <a id="L205"></a>}
    <a id="L206"></a>if e != 0 {
        <a id="L207"></a>return 0, &amp;PathError{&#34;seek&#34;, file.name, Errno(e)}
    <a id="L208"></a>}
    <a id="L209"></a>return r, nil;
<a id="L210"></a>}

<a id="L212"></a><span class="comment">// WriteString is like Write, but writes the contents of string s rather than</span>
<a id="L213"></a><span class="comment">// an array of bytes.</span>
<a id="L214"></a>func (file *File) WriteString(s string) (ret int, err Error) {
    <a id="L215"></a>if file == nil {
        <a id="L216"></a>return 0, EINVAL
    <a id="L217"></a>}
    <a id="L218"></a>b := syscall.StringByteSlice(s);
    <a id="L219"></a>b = b[0 : len(b)-1];
    <a id="L220"></a>return file.Write(b);
<a id="L221"></a>}

<a id="L223"></a><span class="comment">// Pipe returns a connected pair of Files; reads from r return bytes written to w.</span>
<a id="L224"></a><span class="comment">// It returns the files and an Error, if any.</span>
<a id="L225"></a>func Pipe() (r *File, w *File, err Error) {
    <a id="L226"></a>var p [2]int;

    <a id="L228"></a><span class="comment">// See ../syscall/exec.go for description of lock.</span>
    <a id="L229"></a>syscall.ForkLock.RLock();
    <a id="L230"></a>e := syscall.Pipe(&amp;p);
    <a id="L231"></a>if e != 0 {
        <a id="L232"></a>syscall.ForkLock.RUnlock();
        <a id="L233"></a>return nil, nil, NewSyscallError(&#34;pipe&#34;, e);
    <a id="L234"></a>}
    <a id="L235"></a>syscall.CloseOnExec(p[0]);
    <a id="L236"></a>syscall.CloseOnExec(p[1]);
    <a id="L237"></a>syscall.ForkLock.RUnlock();

    <a id="L239"></a>return NewFile(p[0], &#34;|0&#34;), NewFile(p[1], &#34;|1&#34;), nil;
<a id="L240"></a>}

<a id="L242"></a><span class="comment">// Mkdir creates a new directory with the specified name and permission bits.</span>
<a id="L243"></a><span class="comment">// It returns an error, if any.</span>
<a id="L244"></a>func Mkdir(name string, perm int) Error {
    <a id="L245"></a>e := syscall.Mkdir(name, perm);
    <a id="L246"></a>if e != 0 {
        <a id="L247"></a>return &amp;PathError{&#34;mkdir&#34;, name, Errno(e)}
    <a id="L248"></a>}
    <a id="L249"></a>return nil;
<a id="L250"></a>}

<a id="L252"></a><span class="comment">// Stat returns a Dir structure describing the named file and an error, if any.</span>
<a id="L253"></a><span class="comment">// If name names a valid symbolic link, the returned Dir describes</span>
<a id="L254"></a><span class="comment">// the file pointed at by the link and has dir.FollowedSymlink set to true.</span>
<a id="L255"></a><span class="comment">// If name names an invalid symbolic link, the returned Dir describes</span>
<a id="L256"></a><span class="comment">// the link itself and has dir.FollowedSymlink set to false.</span>
<a id="L257"></a>func Stat(name string) (dir *Dir, err Error) {
    <a id="L258"></a>var lstat, stat syscall.Stat_t;
    <a id="L259"></a>e := syscall.Lstat(name, &amp;lstat);
    <a id="L260"></a>if e != 0 {
        <a id="L261"></a>return nil, &amp;PathError{&#34;stat&#34;, name, Errno(e)}
    <a id="L262"></a>}
    <a id="L263"></a>statp := &amp;lstat;
    <a id="L264"></a>if lstat.Mode&amp;syscall.S_IFMT == syscall.S_IFLNK {
        <a id="L265"></a>e := syscall.Stat(name, &amp;stat);
        <a id="L266"></a>if e == 0 {
            <a id="L267"></a>statp = &amp;stat
        <a id="L268"></a>}
    <a id="L269"></a>}
    <a id="L270"></a>return dirFromStat(name, new(Dir), &amp;lstat, statp), nil;
<a id="L271"></a>}

<a id="L273"></a><span class="comment">// Stat returns the Dir structure describing file.</span>
<a id="L274"></a><span class="comment">// It returns the Dir and an error, if any.</span>
<a id="L275"></a>func (file *File) Stat() (dir *Dir, err Error) {
    <a id="L276"></a>var stat syscall.Stat_t;
    <a id="L277"></a>e := syscall.Fstat(file.fd, &amp;stat);
    <a id="L278"></a>if e != 0 {
        <a id="L279"></a>return nil, &amp;PathError{&#34;stat&#34;, file.name, Errno(e)}
    <a id="L280"></a>}
    <a id="L281"></a>return dirFromStat(file.name, new(Dir), &amp;stat, &amp;stat), nil;
<a id="L282"></a>}

<a id="L284"></a><span class="comment">// Lstat returns the Dir structure describing the named file and an error, if any.</span>
<a id="L285"></a><span class="comment">// If the file is a symbolic link, the returned Dir describes the</span>
<a id="L286"></a><span class="comment">// symbolic link.  Lstat makes no attempt to follow the link.</span>
<a id="L287"></a>func Lstat(name string) (dir *Dir, err Error) {
    <a id="L288"></a>var stat syscall.Stat_t;
    <a id="L289"></a>e := syscall.Lstat(name, &amp;stat);
    <a id="L290"></a>if e != 0 {
        <a id="L291"></a>return nil, &amp;PathError{&#34;lstat&#34;, name, Errno(e)}
    <a id="L292"></a>}
    <a id="L293"></a>return dirFromStat(name, new(Dir), &amp;stat, &amp;stat), nil;
<a id="L294"></a>}

<a id="L296"></a><span class="comment">// Readdir reads the contents of the directory associated with file and</span>
<a id="L297"></a><span class="comment">// returns an array of up to count Dir structures, as would be returned</span>
<a id="L298"></a><span class="comment">// by Stat, in directory order.  Subsequent calls on the same file will yield further Dirs.</span>
<a id="L299"></a><span class="comment">// A negative count means to read until EOF.</span>
<a id="L300"></a><span class="comment">// Readdir returns the array and an Error, if any.</span>
<a id="L301"></a>func (file *File) Readdir(count int) (dirs []Dir, err Error) {
    <a id="L302"></a>dirname := file.name;
    <a id="L303"></a>if dirname == &#34;&#34; {
        <a id="L304"></a>dirname = &#34;.&#34;
    <a id="L305"></a>}
    <a id="L306"></a>dirname += &#34;/&#34;;
    <a id="L307"></a>names, err1 := file.Readdirnames(count);
    <a id="L308"></a>if err1 != nil {
        <a id="L309"></a>return nil, err1
    <a id="L310"></a>}
    <a id="L311"></a>dirs = make([]Dir, len(names));
    <a id="L312"></a>for i, filename := range names {
        <a id="L313"></a>dirp, err := Lstat(dirname + filename);
        <a id="L314"></a>if dirp == nil || err != nil {
            <a id="L315"></a>dirs[i].Name = filename <span class="comment">// rest is already zeroed out</span>
        <a id="L316"></a>} else {
            <a id="L317"></a>dirs[i] = *dirp
        <a id="L318"></a>}
    <a id="L319"></a>}
    <a id="L320"></a>return;
<a id="L321"></a>}

<a id="L323"></a><span class="comment">// Chdir changes the current working directory to the named directory.</span>
<a id="L324"></a>func Chdir(dir string) Error {
    <a id="L325"></a>if e := syscall.Chdir(dir); e != 0 {
        <a id="L326"></a>return &amp;PathError{&#34;chdir&#34;, dir, Errno(e)}
    <a id="L327"></a>}
    <a id="L328"></a>return nil;
<a id="L329"></a>}

<a id="L331"></a><span class="comment">// Chdir changes the current working directory to the file,</span>
<a id="L332"></a><span class="comment">// which must be a directory.</span>
<a id="L333"></a>func (f *File) Chdir() Error {
    <a id="L334"></a>if e := syscall.Fchdir(f.fd); e != 0 {
        <a id="L335"></a>return &amp;PathError{&#34;chdir&#34;, f.name, Errno(e)}
    <a id="L336"></a>}
    <a id="L337"></a>return nil;
<a id="L338"></a>}

<a id="L340"></a><span class="comment">// Remove removes the named file or directory.</span>
<a id="L341"></a>func Remove(name string) Error {
    <a id="L342"></a><span class="comment">// System call interface forces us to know</span>
    <a id="L343"></a><span class="comment">// whether name is a file or directory.</span>
    <a id="L344"></a><span class="comment">// Try both: it is cheaper on average than</span>
    <a id="L345"></a><span class="comment">// doing a Stat plus the right one.</span>
    <a id="L346"></a>e := syscall.Unlink(name);
    <a id="L347"></a>if e == 0 {
        <a id="L348"></a>return nil
    <a id="L349"></a>}
    <a id="L350"></a>e1 := syscall.Rmdir(name);
    <a id="L351"></a>if e1 == 0 {
        <a id="L352"></a>return nil
    <a id="L353"></a>}

    <a id="L355"></a><span class="comment">// Both failed: figure out which error to return.</span>
    <a id="L356"></a><span class="comment">// OS X and Linux differ on whether unlink(dir)</span>
    <a id="L357"></a><span class="comment">// returns EISDIR, so can&#39;t use that.  However,</span>
    <a id="L358"></a><span class="comment">// both agree that rmdir(file) returns ENOTDIR,</span>
    <a id="L359"></a><span class="comment">// so we can use that to decide which error is real.</span>
    <a id="L360"></a><span class="comment">// Rmdir might also return ENOTDIR if given a bad</span>
    <a id="L361"></a><span class="comment">// file path, like /etc/passwd/foo, but in that case,</span>
    <a id="L362"></a><span class="comment">// both errors will be ENOTDIR, so it&#39;s okay to</span>
    <a id="L363"></a><span class="comment">// use the error from unlink.</span>
    <a id="L364"></a>if e1 != syscall.ENOTDIR {
        <a id="L365"></a>e = e1
    <a id="L366"></a>}
    <a id="L367"></a>return &amp;PathError{&#34;remove&#34;, name, Errno(e)};
<a id="L368"></a>}

<a id="L370"></a><span class="comment">// LinkError records an error during a link or symlink</span>
<a id="L371"></a><span class="comment">// system call and the paths that caused it.</span>
<a id="L372"></a>type LinkError struct {
    <a id="L373"></a>Op    string;
    <a id="L374"></a>Old   string;
    <a id="L375"></a>New   string;
    <a id="L376"></a>Error Error;
<a id="L377"></a>}

<a id="L379"></a>func (e *LinkError) String() string {
    <a id="L380"></a>return e.Op + &#34; &#34; + e.Old + &#34; &#34; + e.New + &#34;: &#34; + e.Error.String()
<a id="L381"></a>}

<a id="L383"></a><span class="comment">// Link creates a hard link.</span>
<a id="L384"></a>func Link(oldname, newname string) Error {
    <a id="L385"></a>e := syscall.Link(oldname, newname);
    <a id="L386"></a>if e != 0 {
        <a id="L387"></a>return &amp;LinkError{&#34;link&#34;, oldname, newname, Errno(e)}
    <a id="L388"></a>}
    <a id="L389"></a>return nil;
<a id="L390"></a>}

<a id="L392"></a><span class="comment">// Symlink creates a symbolic link.</span>
<a id="L393"></a>func Symlink(oldname, newname string) Error {
    <a id="L394"></a>e := syscall.Symlink(oldname, newname);
    <a id="L395"></a>if e != 0 {
        <a id="L396"></a>return &amp;LinkError{&#34;symlink&#34;, oldname, newname, Errno(e)}
    <a id="L397"></a>}
    <a id="L398"></a>return nil;
<a id="L399"></a>}

<a id="L401"></a><span class="comment">// Readlink reads the contents of a symbolic link: the destination of</span>
<a id="L402"></a><span class="comment">// the link.  It returns the contents and an Error, if any.</span>
<a id="L403"></a>func Readlink(name string) (string, Error) {
    <a id="L404"></a>for len := 128; ; len *= 2 {
        <a id="L405"></a>b := make([]byte, len);
        <a id="L406"></a>n, e := syscall.Readlink(name, b);
        <a id="L407"></a>if e != 0 {
            <a id="L408"></a>return &#34;&#34;, &amp;PathError{&#34;readlink&#34;, name, Errno(e)}
        <a id="L409"></a>}
        <a id="L410"></a>if n &lt; len {
            <a id="L411"></a>return string(b[0:n]), nil
        <a id="L412"></a>}
    <a id="L413"></a>}
    <a id="L414"></a><span class="comment">// Silence 6g.</span>
    <a id="L415"></a>return &#34;&#34;, nil;
<a id="L416"></a>}

<a id="L418"></a><span class="comment">// Chmod changes the mode of the named file to mode.</span>
<a id="L419"></a><span class="comment">// If the file is a symbolic link, it changes the uid and gid of the link&#39;s target.</span>
<a id="L420"></a>func Chmod(name string, mode int) Error {
    <a id="L421"></a>if e := syscall.Chmod(name, mode); e != 0 {
        <a id="L422"></a>return &amp;PathError{&#34;chmod&#34;, name, Errno(e)}
    <a id="L423"></a>}
    <a id="L424"></a>return nil;
<a id="L425"></a>}

<a id="L427"></a><span class="comment">// Chmod changes the mode of the file to mode.</span>
<a id="L428"></a>func (f *File) Chmod(mode int) Error {
    <a id="L429"></a>if e := syscall.Fchmod(f.fd, mode); e != 0 {
        <a id="L430"></a>return &amp;PathError{&#34;chmod&#34;, f.name, Errno(e)}
    <a id="L431"></a>}
    <a id="L432"></a>return nil;
<a id="L433"></a>}

<a id="L435"></a><span class="comment">// Chown changes the numeric uid and gid of the named file.</span>
<a id="L436"></a><span class="comment">// If the file is a symbolic link, it changes the uid and gid of the link&#39;s target.</span>
<a id="L437"></a>func Chown(name string, uid, gid int) Error {
    <a id="L438"></a>if e := syscall.Chown(name, uid, gid); e != 0 {
        <a id="L439"></a>return &amp;PathError{&#34;chown&#34;, name, Errno(e)}
    <a id="L440"></a>}
    <a id="L441"></a>return nil;
<a id="L442"></a>}

<a id="L444"></a><span class="comment">// Lchown changes the numeric uid and gid of the named file.</span>
<a id="L445"></a><span class="comment">// If the file is a symbolic link, it changes the uid and gid of the link itself.</span>
<a id="L446"></a>func Lchown(name string, uid, gid int) Error {
    <a id="L447"></a>if e := syscall.Lchown(name, uid, gid); e != 0 {
        <a id="L448"></a>return &amp;PathError{&#34;lchown&#34;, name, Errno(e)}
    <a id="L449"></a>}
    <a id="L450"></a>return nil;
<a id="L451"></a>}

<a id="L453"></a><span class="comment">// Chown changes the numeric uid and gid of the named file.</span>
<a id="L454"></a>func (f *File) Chown(uid, gid int) Error {
    <a id="L455"></a>if e := syscall.Fchown(f.fd, uid, gid); e != 0 {
        <a id="L456"></a>return &amp;PathError{&#34;chown&#34;, f.name, Errno(e)}
    <a id="L457"></a>}
    <a id="L458"></a>return nil;
<a id="L459"></a>}

<a id="L461"></a><span class="comment">// Truncate changes the size of the named file.</span>
<a id="L462"></a><span class="comment">// If the file is a symbolic link, it changes the size of the link&#39;s target.</span>
<a id="L463"></a>func Truncate(name string, size int64) Error {
    <a id="L464"></a>if e := syscall.Truncate(name, size); e != 0 {
        <a id="L465"></a>return &amp;PathError{&#34;truncate&#34;, name, Errno(e)}
    <a id="L466"></a>}
    <a id="L467"></a>return nil;
<a id="L468"></a>}

<a id="L470"></a><span class="comment">// Truncate changes the size of the file.</span>
<a id="L471"></a><span class="comment">// It does not change the I/O offset.</span>
<a id="L472"></a>func (f *File) Truncate(size int64) Error {
    <a id="L473"></a>if e := syscall.Ftruncate(f.fd, size); e != 0 {
        <a id="L474"></a>return &amp;PathError{&#34;truncate&#34;, f.name, Errno(e)}
    <a id="L475"></a>}
    <a id="L476"></a>return nil;
<a id="L477"></a>}
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

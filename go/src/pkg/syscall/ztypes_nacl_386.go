<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/ztypes_nacl_386.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/syscall/ztypes_nacl_386.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// godefs -gsyscall -f-m32 -f-I/home/rsc/pub/nacl/native_client/src/third_party/nacl_sdk/linux/sdk/nacl-sdk/nacl/include -f-I/home/rsc/pub/nacl/native_client types_nacl.c</span>

<a id="L3"></a><span class="comment">// MACHINE GENERATED - DO NOT EDIT.</span>

<a id="L5"></a>package syscall

<a id="L7"></a><span class="comment">// Constants</span>
<a id="L8"></a>const (
    <a id="L9"></a>sizeofPtr      = 0x4;
    <a id="L10"></a>sizeofShort    = 0x2;
    <a id="L11"></a>sizeofInt      = 0x4;
    <a id="L12"></a>sizeofLong     = 0x4;
    <a id="L13"></a>sizeofLongLong = 0x8;
    <a id="L14"></a>PROT_READ      = 0x1;
    <a id="L15"></a>PROT_WRITE     = 0x2;
    <a id="L16"></a>MAP_SHARED     = 0x1;
    <a id="L17"></a>SYS_FORK       = 0;
    <a id="L18"></a>SYS_PTRACE     = 0;
    <a id="L19"></a>SYS_CHDIR      = 0;
    <a id="L20"></a>SYS_DUP2       = 0;
    <a id="L21"></a>SYS_FCNTL      = 0;
    <a id="L22"></a>SYS_EXECVE     = 0;
    <a id="L23"></a>O_RDONLY       = 0;
    <a id="L24"></a>O_WRONLY       = 0x1;
    <a id="L25"></a>O_RDWR         = 0x2;
    <a id="L26"></a>O_APPEND       = 0x400;
    <a id="L27"></a>O_ASYNC        = 0x2000;
    <a id="L28"></a>O_CREAT        = 0x40;
    <a id="L29"></a>O_NOCTTY       = 0;
    <a id="L30"></a>O_NONBLOCK     = 0x800;
    <a id="L31"></a>O_SYNC         = 0x1000;
    <a id="L32"></a>O_TRUNC        = 0x200;
    <a id="L33"></a>O_CLOEXEC      = 0;
    <a id="L34"></a>F_GETFD        = 0x1;
    <a id="L35"></a>F_SETFD        = 0x2;
    <a id="L36"></a>F_GETFL        = 0x3;
    <a id="L37"></a>F_SETFL        = 0x4;
    <a id="L38"></a>FD_CLOEXEC     = 0;
    <a id="L39"></a>S_IFMT         = 0x1f000;
    <a id="L40"></a>S_IFIFO        = 0x1000;
    <a id="L41"></a>S_IFCHR        = 0x2000;
    <a id="L42"></a>S_IFDIR        = 0x4000;
    <a id="L43"></a>S_IFBLK        = 0x6000;
    <a id="L44"></a>S_IFREG        = 0x8000;
    <a id="L45"></a>S_IFLNK        = 0xa000;
    <a id="L46"></a>S_IFSOCK       = 0xc000;
    <a id="L47"></a>S_ISUID        = 0x800;
    <a id="L48"></a>S_ISGID        = 0x400;
    <a id="L49"></a>S_ISVTX        = 0x200;
    <a id="L50"></a>S_IRUSR        = 0x100;
    <a id="L51"></a>S_IWUSR        = 0x80;
    <a id="L52"></a>S_IXUSR        = 0x40;
<a id="L53"></a>)

<a id="L55"></a><span class="comment">// Types</span>

<a id="L57"></a>type _C_short int16

<a id="L59"></a>type _C_int int32

<a id="L61"></a>type _C_long int32

<a id="L63"></a>type _C_long_long int64

<a id="L65"></a>type _C_off_t int32

<a id="L67"></a>type Timespec struct {
    <a id="L68"></a>Sec  int32;
    <a id="L69"></a>Nsec int32;
<a id="L70"></a>}

<a id="L72"></a>type Timeval struct {
    <a id="L73"></a>Sec  int32;
    <a id="L74"></a>Usec int32;
<a id="L75"></a>}

<a id="L77"></a>type Time_t int32

<a id="L79"></a>type _Gid_t uint32

<a id="L81"></a>type Stat_t struct {
    <a id="L82"></a>Dev       int64;
    <a id="L83"></a>Ino       uint32;
    <a id="L84"></a>Mode      uint32;
    <a id="L85"></a>Nlink     uint32;
    <a id="L86"></a>Uid       uint32;
    <a id="L87"></a>Gid       uint32;
    <a id="L88"></a>__padding int32;
    <a id="L89"></a>Rdev      int64;
    <a id="L90"></a>Size      int32;
    <a id="L91"></a>Blksize   int32;
    <a id="L92"></a>Blocks    int32;
    <a id="L93"></a>Atime     int32;
    <a id="L94"></a>Mtime     int32;
    <a id="L95"></a>Ctime     int32;
<a id="L96"></a>}

<a id="L98"></a>type Dirent struct {
    <a id="L99"></a>Ino    uint32;
    <a id="L100"></a>Off    int32;
    <a id="L101"></a>Reclen uint16;
    <a id="L102"></a>Name   [256]int8;
    <a id="L103"></a>Pad0   [2]byte;
<a id="L104"></a>}
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

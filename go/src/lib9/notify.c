<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/notify.c</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/lib9/notify.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
/*
Plan 9 from User Space src/lib9/notify.c
http://code.swtch.com/plan9port/src/tip/src/lib9/notify.c

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

/*
 * Signal handling for Plan 9 programs.
 * We stubbornly use the strings from Plan 9 instead
 * of the enumerated Unix constants.
 * There are some weird translations.  In particular,
 * a &#34;kill&#34; note is the same as SIGTERM in Unix.
 * There is no equivalent note to Unix&#39;s SIGKILL, since
 * it&#39;s not a deliverable signal anyway.
 *
 * We do not handle SIGABRT or SIGSEGV, mainly because
 * the thread library queues its notes for later, and we want
 * to dump core with the state at time of delivery.
 *
 * We have to add some extra entry points to provide the
 * ability to tweak which signals are deliverable and which
 * are acted upon.  Notifydisable and notifyenable play with
 * the process signal mask.  Notifyignore enables the signal
 * but will not call notifyf when it comes in.  This is occasionally
 * useful.
 */

#include &lt;u.h&gt;
#include &lt;signal.h&gt;
#define NOPLAN9DEFINES
#include &lt;libc.h&gt;

extern char *_p9sigstr(int, char*);
extern int _p9strsig(char*);

typedef struct Sig Sig;
struct Sig
{
	int sig;			/* signal number */
	int flags;
};

enum
{
	Restart = 1&lt;&lt;0,
	Ignore = 1&lt;&lt;1
};

static Sig sigs[] = {
	SIGHUP,		0,
	SIGINT,		0,
	SIGQUIT,		0,
	SIGILL,		0,
	SIGTRAP,		0,
/*	SIGABRT, 		0, 	*/
#ifdef SIGEMT
	SIGEMT,		0,
#endif
	SIGFPE,		0,
	SIGBUS,		0,
/*	SIGSEGV, 		0, 	*/
	SIGCHLD,		Restart|Ignore,
	SIGSYS,		0,
	SIGPIPE,		Ignore,
	SIGALRM,		0,
	SIGTERM,		0,
	SIGTSTP,		Restart|Ignore,
/*	SIGTTIN,		Restart|Ignore, */
/*	SIGTTOU,		Restart|Ignore, */
	SIGXCPU,		0,
	SIGXFSZ,		0,
	SIGVTALRM,	0,
	SIGUSR1,		0,
	SIGUSR2,		0,
#ifdef SIGWINCH
	SIGWINCH,	Restart|Ignore,
#endif
#ifdef SIGINFO
	SIGINFO,		Restart|Ignore,
#endif
};

static Sig*
findsig(int s)
{
	int i;

	for(i=0; i&lt;nelem(sigs); i++)
		if(sigs[i].sig == s)
			return &amp;sigs[i];
	return nil;
}

/*
 * The thread library initializes _notejmpbuf to its own
 * routine which provides a per-pthread jump buffer.
 * If we&#39;re not using the thread library, we assume we are
 * single-threaded.
 */
typedef struct Jmp Jmp;
struct Jmp
{
	p9jmp_buf b;
};

static Jmp onejmp;

static Jmp*
getonejmp(void)
{
	return &amp;onejmp;
}

Jmp *(*_notejmpbuf)(void) = getonejmp;
static void noteinit(void);

/*
 * Actual signal handler.
 */

static void (*notifyf)(void*, char*);	/* Plan 9 handler */

static void
signotify(int sig)
{
	char tmp[64];
	Jmp *j;
	Sig *s;

	j = (*_notejmpbuf)();
	switch(p9setjmp(j-&gt;b)){
	case 0:
		if(notifyf)
			(*notifyf)(nil, _p9sigstr(sig, tmp));
		/* fall through */
	case 1:	/* noted(NDFLT) */
		if(0)print(&#34;DEFAULT %d\n&#34;, sig);
		s = findsig(sig);
		if(s &amp;&amp; (s-&gt;flags&amp;Ignore))
			return;
		signal(sig, SIG_DFL);
		raise(sig);
		_exit(1);
	case 2:	/* noted(NCONT) */
		if(0)print(&#34;HANDLED %d\n&#34;, sig);
		return;
	}
}

static void
signonotify(int sig)
{
	USED(sig);
}

int
noted(int v)
{
	p9longjmp((*_notejmpbuf)()-&gt;b, v==NCONT ? 2 : 1);
	abort();
	return 0;
}

int
notify(void (*f)(void*, char*))
{
	static int init;

	notifyf = f;
	if(!init){
		init = 1;
		noteinit();
	}
	return 0;
}

/*
 * Nonsense about enabling and disabling signals.
 */
typedef void Sighandler(int);
static Sighandler*
handler(int s)
{
	struct sigaction sa;

	sigaction(s, nil, &amp;sa);
	return sa.sa_handler;
}

static int
notesetenable(int sig, int enabled)
{
	sigset_t mask, omask;

	if(sig == 0)
		return -1;

	sigemptyset(&amp;mask);
	sigaddset(&amp;mask, sig);
	sigprocmask(enabled ? SIG_UNBLOCK : SIG_BLOCK, &amp;mask, &amp;omask);
	return !sigismember(&amp;omask, sig);
}

int
noteenable(char *msg)
{
	return notesetenable(_p9strsig(msg), 1);
}

int
notedisable(char *msg)
{
	return notesetenable(_p9strsig(msg), 0);
}

static int
notifyseton(int s, int on)
{
	Sig *sig;
	struct sigaction sa, osa;

	sig = findsig(s);
	if(sig == nil)
		return -1;
	memset(&amp;sa, 0, sizeof sa);
	sa.sa_handler = on ? signotify : signonotify;
	if(sig-&gt;flags&amp;Restart)
		sa.sa_flags |= SA_RESTART;

	/*
	 * We can&#39;t allow signals within signals because there&#39;s
	 * only one jump buffer.
	 */
	sigfillset(&amp;sa.sa_mask);

	/*
	 * Install handler.
	 */
	sigaction(sig-&gt;sig, &amp;sa, &amp;osa);
	return osa.sa_handler == signotify;
}

int
notifyon(char *msg)
{
	return notifyseton(_p9strsig(msg), 1);
}

int
notifyoff(char *msg)
{
	return notifyseton(_p9strsig(msg), 0);
}

/*
 * Initialization follows sigs table.
 */
static void
noteinit(void)
{
	int i;
	Sig *sig;

	for(i=0; i&lt;nelem(sigs); i++){
		sig = &amp;sigs[i];
		/*
		 * If someone has already installed a handler,
		 * It&#39;s probably some ld preload nonsense,
		 * like pct (a SIGVTALRM-based profiler).
		 * Or maybe someone has already called notifyon/notifyoff.
		 * Leave it alone.
		 */
		if(handler(sig-&gt;sig) != SIG_DFL)
			continue;
		notifyseton(sig-&gt;sig, 1);
	}
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

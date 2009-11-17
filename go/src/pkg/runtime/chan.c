<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/chan.c</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/chan.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;
#include &#34;type.h&#34;

static	int32	debug	= 0;
static	Lock		chanlock;

enum
{
	Wclosed		= 0x0001,	// writer has closed
	Rclosed		= 0x0002,	// reader has seen close
	Eincr		= 0x0004,	// increment errors
	Emax		= 0x0800,	// error limit before throw
};

typedef	struct	Link	Link;
typedef	struct	WaitQ	WaitQ;
typedef	struct	SudoG	SudoG;
typedef	struct	Select	Select;
typedef	struct	Scase	Scase;

struct	SudoG
{
	G*	g;		// g and selgen constitute
	int32	selgen;		// a weak pointer to g
	int16	offset;		// offset of case number
	int8	isfree;		// offset of case number
	SudoG*	link;
	byte	elem[8];	// synch data element (+ more)
};

struct	WaitQ
{
	SudoG*	first;
	SudoG*	last;
};

struct	Hchan
{
	uint32	qcount;			// total data in the q
	uint32	dataqsiz;		// size of the circular q
	uint16	elemsize;
	uint16	closed;			// Wclosed Rclosed errorcount
	uint8	elemalign;
	Alg*	elemalg;		// interface for element type
	Link*	senddataq;		// pointer for sender
	Link*	recvdataq;		// pointer for receiver
	WaitQ	recvq;			// list of recv waiters
	WaitQ	sendq;			// list of send waiters
	SudoG*	free;			// freelist
};

struct	Link
{
	Link*	link;			// asynch queue circular linked list
	byte	elem[8];		// asynch queue data element (+ more)
};

struct	Scase
{
	Hchan*	chan;			// chan
	byte*	pc;			// return pc
	uint16	send;			// 0-recv 1-send 2-default
	uint16	so;			// vararg of selected bool
	union {
		byte	elem[8];	// element (send)
		byte*	elemp;		// pointer to element (recv)
	} u;
};

struct	Select
{
	uint16	tcase;			// total count of scase[]
	uint16	ncase;			// currently filled scase[]
	Select*	link;			// for freelist
	Scase*	scase[1];		// one per case
};

static	SudoG*	dequeue(WaitQ*, Hchan*);
static	void	enqueue(WaitQ*, SudoG*);
static	SudoG*	allocsg(Hchan*);
static	void	freesg(Hchan*, SudoG*);
static	uint32	gcd(uint32, uint32);
static	uint32	fastrand1(void);
static	uint32	fastrand2(void);

Hchan*
makechan(Type *elem, uint32 hint)
{
	Hchan *c;
	int32 i;

	if(elem-&gt;alg &gt;= nelem(algarray)) {
		printf(&#34;chan(alg=%d)\n&#34;, elem-&gt;alg);
		throw(&#34;runtime·makechan: unsupported elem type&#34;);
	}

	c = mal(sizeof(*c));

	c-&gt;elemsize = elem-&gt;size;
	c-&gt;elemalg = &amp;algarray[elem-&gt;alg];
	c-&gt;elemalign = elem-&gt;align;

	if(hint &gt; 0) {
		Link *d, *b, *e;

		// make a circular q
		b = nil;
		e = nil;
		for(i=0; i&lt;hint; i++) {
			d = mal(sizeof(*d) + c-&gt;elemsize - sizeof(d-&gt;elem));
			if(e == nil)
				e = d;
			d-&gt;link = b;
			b = d;
		}
		e-&gt;link = b;
		c-&gt;recvdataq = b;
		c-&gt;senddataq = b;
		c-&gt;qcount = 0;
		c-&gt;dataqsiz = hint;
	}

	if(debug) {
		prints(&#34;makechan: chan=&#34;);
		runtime·printpointer(c);
		prints(&#34;; elemsize=&#34;);
		runtime·printint(elem-&gt;size);
		prints(&#34;; elemalg=&#34;);
		runtime·printint(elem-&gt;alg);
		prints(&#34;; elemalign=&#34;);
		runtime·printint(elem-&gt;align);
		prints(&#34;; dataqsiz=&#34;);
		runtime·printint(c-&gt;dataqsiz);
		prints(&#34;\n&#34;);
	}

	return c;
}

// makechan(elemsize uint32, elemalg uint32, hint uint32) (hchan *chan any);
void
runtime·makechan(Type *elem, uint32 hint, Hchan *ret)
{
	ret = makechan(elem, hint);
	FLUSH(&amp;ret);
}

static void
incerr(Hchan* c)
{
	c-&gt;closed += Eincr;
	if(c-&gt;closed &amp; Emax) {
		unlock(&amp;chanlock);
		throw(&#34;too many operations on a closed channel&#34;);
	}
}

/*
 * generic single channel send/recv
 * if the bool pointer is nil,
 * then the full exchange will
 * occur. if pres is not nil,
 * then the protocol will not
 * sleep but return if it could
 * not complete
 */
void
chansend(Hchan *c, byte *ep, bool *pres)
{
	SudoG *sg;
	G* gp;

	if(debug) {
		prints(&#34;chansend: chan=&#34;);
		runtime·printpointer(c);
		prints(&#34;; elem=&#34;);
		c-&gt;elemalg-&gt;print(c-&gt;elemsize, ep);
		prints(&#34;\n&#34;);
	}

	lock(&amp;chanlock);
loop:
	if(c-&gt;closed &amp; Wclosed)
		goto closed;

	if(c-&gt;dataqsiz &gt; 0)
		goto asynch;

	sg = dequeue(&amp;c-&gt;recvq, c);
	if(sg != nil) {
		if(ep != nil)
			c-&gt;elemalg-&gt;copy(c-&gt;elemsize, sg-&gt;elem, ep);

		gp = sg-&gt;g;
		gp-&gt;param = sg;
		unlock(&amp;chanlock);
		ready(gp);

		if(pres != nil)
			*pres = true;
		return;
	}

	if(pres != nil) {
		unlock(&amp;chanlock);
		*pres = false;
		return;
	}

	sg = allocsg(c);
	if(ep != nil)
		c-&gt;elemalg-&gt;copy(c-&gt;elemsize, sg-&gt;elem, ep);
	g-&gt;param = nil;
	g-&gt;status = Gwaiting;
	enqueue(&amp;c-&gt;sendq, sg);
	unlock(&amp;chanlock);
	gosched();

	lock(&amp;chanlock);
	sg = g-&gt;param;
	if(sg == nil)
		goto loop;
	freesg(c, sg);
	unlock(&amp;chanlock);
	return;

asynch:
	if(c-&gt;closed &amp; Wclosed)
		goto closed;

	if(c-&gt;qcount &gt;= c-&gt;dataqsiz) {
		if(pres != nil) {
			unlock(&amp;chanlock);
			*pres = false;
			return;
		}
		sg = allocsg(c);
		g-&gt;status = Gwaiting;
		enqueue(&amp;c-&gt;sendq, sg);
		unlock(&amp;chanlock);
		gosched();

		lock(&amp;chanlock);
		goto asynch;
	}
	if(ep != nil)
		c-&gt;elemalg-&gt;copy(c-&gt;elemsize, c-&gt;senddataq-&gt;elem, ep);
	c-&gt;senddataq = c-&gt;senddataq-&gt;link;
	c-&gt;qcount++;

	sg = dequeue(&amp;c-&gt;recvq, c);
	if(sg != nil) {
		gp = sg-&gt;g;
		freesg(c, sg);
		unlock(&amp;chanlock);
		ready(gp);
	} else
		unlock(&amp;chanlock);
	if(pres != nil)
		*pres = true;
	return;

closed:
	incerr(c);
	if(pres != nil)
		*pres = true;
	unlock(&amp;chanlock);
}

void
chanrecv(Hchan* c, byte *ep, bool* pres)
{
	SudoG *sg;
	G *gp;

	if(debug) {
		prints(&#34;chanrecv: chan=&#34;);
		runtime·printpointer(c);
		prints(&#34;\n&#34;);
	}

	lock(&amp;chanlock);
loop:
	if(c-&gt;dataqsiz &gt; 0)
		goto asynch;

	if(c-&gt;closed &amp; Wclosed)
		goto closed;

	sg = dequeue(&amp;c-&gt;sendq, c);
	if(sg != nil) {
		c-&gt;elemalg-&gt;copy(c-&gt;elemsize, ep, sg-&gt;elem);

		gp = sg-&gt;g;
		gp-&gt;param = sg;
		unlock(&amp;chanlock);
		ready(gp);

		if(pres != nil)
			*pres = true;
		return;
	}

	if(pres != nil) {
		unlock(&amp;chanlock);
		*pres = false;
		return;
	}

	sg = allocsg(c);
	g-&gt;param = nil;
	g-&gt;status = Gwaiting;
	enqueue(&amp;c-&gt;recvq, sg);
	unlock(&amp;chanlock);
	gosched();

	lock(&amp;chanlock);
	sg = g-&gt;param;
	if(sg == nil)
		goto loop;

	c-&gt;elemalg-&gt;copy(c-&gt;elemsize, ep, sg-&gt;elem);
	freesg(c, sg);
	unlock(&amp;chanlock);
	return;

asynch:
	if(c-&gt;qcount &lt;= 0) {
		if(c-&gt;closed &amp; Wclosed)
			goto closed;

		if(pres != nil) {
			unlock(&amp;chanlock);
			*pres = false;
			return;
		}
		sg = allocsg(c);
		g-&gt;status = Gwaiting;
		enqueue(&amp;c-&gt;recvq, sg);
		unlock(&amp;chanlock);
		gosched();

		lock(&amp;chanlock);
		goto asynch;
	}
	c-&gt;elemalg-&gt;copy(c-&gt;elemsize, ep, c-&gt;recvdataq-&gt;elem);
	c-&gt;recvdataq = c-&gt;recvdataq-&gt;link;
	c-&gt;qcount--;
	sg = dequeue(&amp;c-&gt;sendq, c);
	if(sg != nil) {
		gp = sg-&gt;g;
		freesg(c, sg);
		unlock(&amp;chanlock);
		ready(gp);
		if(pres != nil)
			*pres = true;
		return;
	}

	unlock(&amp;chanlock);
	if(pres != nil)
		*pres = true;
	return;

closed:
	c-&gt;elemalg-&gt;copy(c-&gt;elemsize, ep, nil);
	c-&gt;closed |= Rclosed;
	incerr(c);
	if(pres != nil)
		*pres = true;
	unlock(&amp;chanlock);
}

// chansend1(hchan *chan any, elem any);
void
runtime·chansend1(Hchan* c, ...)
{
	int32 o;
	byte *ae;

	o = rnd(sizeof(c), c-&gt;elemalign);
	ae = (byte*)&amp;c + o;
	chansend(c, ae, nil);
}

// chansend2(hchan *chan any, elem any) (pres bool);
void
runtime·chansend2(Hchan* c, ...)
{
	int32 o;
	byte *ae, *ap;

	o = rnd(sizeof(c), c-&gt;elemalign);
	ae = (byte*)&amp;c + o;
	o = rnd(o+c-&gt;elemsize, Structrnd);
	ap = (byte*)&amp;c + o;

	chansend(c, ae, ap);
}

// chanrecv1(hchan *chan any) (elem any);
void
runtime·chanrecv1(Hchan* c, ...)
{
	int32 o;
	byte *ae;

	o = rnd(sizeof(c), Structrnd);
	ae = (byte*)&amp;c + o;

	chanrecv(c, ae, nil);
}

// chanrecv2(hchan *chan any) (elem any, pres bool);
void
runtime·chanrecv2(Hchan* c, ...)
{
	int32 o;
	byte *ae, *ap;

	o = rnd(sizeof(c), Structrnd);
	ae = (byte*)&amp;c + o;
	o = rnd(o+c-&gt;elemsize, 1);
	ap = (byte*)&amp;c + o;

	chanrecv(c, ae, ap);
}

// newselect(size uint32) (sel *byte);
void
runtime·newselect(int32 size, ...)
{
	int32 n, o;
	Select **selp;
	Select *sel;

	o = rnd(sizeof(size), Structrnd);
	selp = (Select**)((byte*)&amp;size + o);
	n = 0;
	if(size &gt; 1)
		n = size-1;

	sel = mal(sizeof(*sel) + n*sizeof(sel-&gt;scase[0]));

	sel-&gt;tcase = size;
	sel-&gt;ncase = 0;
	*selp = sel;
	if(debug) {
		prints(&#34;newselect s=&#34;);
		runtime·printpointer(sel);
		prints(&#34; size=&#34;);
		runtime·printint(size);
		prints(&#34;\n&#34;);
	}
}

// selectsend(sel *byte, hchan *chan any, elem any) (selected bool);
void
runtime·selectsend(Select *sel, Hchan *c, ...)
{
	int32 i, eo;
	Scase *cas;
	byte *ae;

	// nil cases do not compete
	if(c == nil)
		return;

	i = sel-&gt;ncase;
	if(i &gt;= sel-&gt;tcase)
		throw(&#34;selectsend: too many cases&#34;);
	sel-&gt;ncase = i+1;
	cas = mal(sizeof *cas + c-&gt;elemsize - sizeof(cas-&gt;u.elem));
	sel-&gt;scase[i] = cas;

	cas-&gt;pc = runtime·getcallerpc(&amp;sel);
	cas-&gt;chan = c;

	eo = rnd(sizeof(sel), sizeof(c));
	eo = rnd(eo+sizeof(c), c-&gt;elemsize);
	cas-&gt;so = rnd(eo+c-&gt;elemsize, Structrnd);
	cas-&gt;send = 1;

	ae = (byte*)&amp;sel + eo;
	c-&gt;elemalg-&gt;copy(c-&gt;elemsize, cas-&gt;u.elem, ae);

	if(debug) {
		prints(&#34;selectsend s=&#34;);
		runtime·printpointer(sel);
		prints(&#34; pc=&#34;);
		runtime·printpointer(cas-&gt;pc);
		prints(&#34; chan=&#34;);
		runtime·printpointer(cas-&gt;chan);
		prints(&#34; so=&#34;);
		runtime·printint(cas-&gt;so);
		prints(&#34; send=&#34;);
		runtime·printint(cas-&gt;send);
		prints(&#34;\n&#34;);
	}
}

// selectrecv(sel *byte, hchan *chan any, elem *any) (selected bool);
void
runtime·selectrecv(Select *sel, Hchan *c, ...)
{
	int32 i, eo;
	Scase *cas;

	// nil cases do not compete
	if(c == nil)
		return;

	i = sel-&gt;ncase;
	if(i &gt;= sel-&gt;tcase)
		throw(&#34;selectrecv: too many cases&#34;);
	sel-&gt;ncase = i+1;
	cas = mal(sizeof *cas);
	sel-&gt;scase[i] = cas;
	cas-&gt;pc = runtime·getcallerpc(&amp;sel);
	cas-&gt;chan = c;

	eo = rnd(sizeof(sel), sizeof(c));
	eo = rnd(eo+sizeof(c), sizeof(byte*));
	cas-&gt;so = rnd(eo+sizeof(byte*), Structrnd);
	cas-&gt;send = 0;
	cas-&gt;u.elemp = *(byte**)((byte*)&amp;sel + eo);

	if(debug) {
		prints(&#34;selectrecv s=&#34;);
		runtime·printpointer(sel);
		prints(&#34; pc=&#34;);
		runtime·printpointer(cas-&gt;pc);
		prints(&#34; chan=&#34;);
		runtime·printpointer(cas-&gt;chan);
		prints(&#34; so=&#34;);
		runtime·printint(cas-&gt;so);
		prints(&#34; send=&#34;);
		runtime·printint(cas-&gt;send);
		prints(&#34;\n&#34;);
	}
}


// selectdefaul(sel *byte) (selected bool);
void
runtime·selectdefault(Select *sel, ...)
{
	int32 i;
	Scase *cas;

	i = sel-&gt;ncase;
	if(i &gt;= sel-&gt;tcase)
		throw(&#34;selectdefault: too many cases&#34;);
	sel-&gt;ncase = i+1;
	cas = mal(sizeof *cas);
	sel-&gt;scase[i] = cas;
	cas-&gt;pc = runtime·getcallerpc(&amp;sel);
	cas-&gt;chan = nil;

	cas-&gt;so = rnd(sizeof(sel), Structrnd);
	cas-&gt;send = 2;
	cas-&gt;u.elemp = nil;

	if(debug) {
		prints(&#34;selectdefault s=&#34;);
		runtime·printpointer(sel);
		prints(&#34; pc=&#34;);
		runtime·printpointer(cas-&gt;pc);
		prints(&#34; so=&#34;);
		runtime·printint(cas-&gt;so);
		prints(&#34; send=&#34;);
		runtime·printint(cas-&gt;send);
		prints(&#34;\n&#34;);
	}
}

static void
freesel(Select *sel)
{
	uint32 i;

	for(i=0; i&lt;sel-&gt;ncase; i++)
		free(sel-&gt;scase[i]);
	free(sel);
}

// selectgo(sel *byte);
void
runtime·selectgo(Select *sel)
{
	uint32 p, o, i;
	Scase *cas, *dfl;
	Hchan *c;
	SudoG *sg;
	G *gp;
	byte *as;

	if(debug) {
		prints(&#34;selectgo: sel=&#34;);
		runtime·printpointer(sel);
		prints(&#34;\n&#34;);
	}

	if(sel-&gt;ncase &lt; 2) {
		if(sel-&gt;ncase &lt; 1)
			throw(&#34;selectgo: no cases&#34;);
		// make special case of one.
	}

	// select a (relative) prime
	for(i=0;; i++) {
		p = fastrand1();
		if(gcd(p, sel-&gt;ncase) == 1)
			break;
		if(i &gt; 1000) {
			throw(&#34;selectgo: failed to select prime&#34;);
		}
	}

	// select an initial offset
	o = fastrand2();

	p %= sel-&gt;ncase;
	o %= sel-&gt;ncase;

	lock(&amp;chanlock);

loop:
	// pass 1 - look for something already waiting
	dfl = nil;
	for(i=0; i&lt;sel-&gt;ncase; i++) {
		cas = sel-&gt;scase[o];

		if(cas-&gt;send == 2) {	// default
			dfl = cas;
			goto next1;
		}

		c = cas-&gt;chan;
		if(c-&gt;dataqsiz &gt; 0) {
			if(cas-&gt;send) {
				if(c-&gt;closed &amp; Wclosed)
					goto sclose;
				if(c-&gt;qcount &lt; c-&gt;dataqsiz)
					goto asyns;
				goto next1;
			}
			if(c-&gt;qcount &gt; 0)
				goto asynr;
			if(c-&gt;closed &amp; Wclosed)
				goto rclose;
			goto next1;
		}

		if(cas-&gt;send) {
			if(c-&gt;closed &amp; Wclosed)
				goto sclose;
			sg = dequeue(&amp;c-&gt;recvq, c);
			if(sg != nil)
				goto gots;
			goto next1;
		}
		sg = dequeue(&amp;c-&gt;sendq, c);
		if(sg != nil)
			goto gotr;
		if(c-&gt;closed &amp; Wclosed)
			goto rclose;

	next1:
		o += p;
		if(o &gt;= sel-&gt;ncase)
			o -= sel-&gt;ncase;
	}

	if(dfl != nil) {
		cas = dfl;
		goto retc;
	}


	// pass 2 - enqueue on all chans
	for(i=0; i&lt;sel-&gt;ncase; i++) {
		cas = sel-&gt;scase[o];
		c = cas-&gt;chan;

		if(c-&gt;dataqsiz &gt; 0) {
			if(cas-&gt;send) {
				if(c-&gt;qcount &lt; c-&gt;dataqsiz) {
					prints(&#34;selectgo: pass 2 async send\n&#34;);
					goto asyns;
				}
				sg = allocsg(c);
				sg-&gt;offset = o;
				enqueue(&amp;c-&gt;sendq, sg);
				goto next2;
			}
			if(c-&gt;qcount &gt; 0) {
				prints(&#34;selectgo: pass 2 async recv\n&#34;);
				goto asynr;
			}
			sg = allocsg(c);
			sg-&gt;offset = o;
			enqueue(&amp;c-&gt;recvq, sg);
			goto next2;
		}

		if(cas-&gt;send) {
			sg = dequeue(&amp;c-&gt;recvq, c);
			if(sg != nil) {
				prints(&#34;selectgo: pass 2 sync send\n&#34;);
				g-&gt;selgen++;
				goto gots;
			}
			sg = allocsg(c);
			sg-&gt;offset = o;
			c-&gt;elemalg-&gt;copy(c-&gt;elemsize, sg-&gt;elem, cas-&gt;u.elem);
			enqueue(&amp;c-&gt;sendq, sg);
			goto next2;
		}
		sg = dequeue(&amp;c-&gt;sendq, c);
		if(sg != nil) {
			prints(&#34;selectgo: pass 2 sync recv\n&#34;);
			g-&gt;selgen++;
			goto gotr;
		}
		sg = allocsg(c);
		sg-&gt;offset = o;
		enqueue(&amp;c-&gt;recvq, sg);

	next2:
		o += p;
		if(o &gt;= sel-&gt;ncase)
			o -= sel-&gt;ncase;
	}

	g-&gt;param = nil;
	g-&gt;status = Gwaiting;
	unlock(&amp;chanlock);
	gosched();

	lock(&amp;chanlock);
	sg = g-&gt;param;
	if(sg == nil)
		goto loop;

	o = sg-&gt;offset;
	cas = sel-&gt;scase[o];
	c = cas-&gt;chan;

	if(c-&gt;dataqsiz &gt; 0) {
//		prints(&#34;shouldnt happen\n&#34;);
		goto loop;
	}

	if(debug) {
		prints(&#34;wait-return: sel=&#34;);
		runtime·printpointer(sel);
		prints(&#34; c=&#34;);
		runtime·printpointer(c);
		prints(&#34; cas=&#34;);
		runtime·printpointer(cas);
		prints(&#34; send=&#34;);
		runtime·printint(cas-&gt;send);
		prints(&#34; o=&#34;);
		runtime·printint(o);
		prints(&#34;\n&#34;);
	}

	if(!cas-&gt;send) {
		if(cas-&gt;u.elemp != nil)
			c-&gt;elemalg-&gt;copy(c-&gt;elemsize, cas-&gt;u.elemp, sg-&gt;elem);
	}

	freesg(c, sg);
	goto retc;

asynr:
	if(cas-&gt;u.elemp != nil)
		c-&gt;elemalg-&gt;copy(c-&gt;elemsize, cas-&gt;u.elemp, c-&gt;recvdataq-&gt;elem);
	c-&gt;recvdataq = c-&gt;recvdataq-&gt;link;
	c-&gt;qcount--;
	sg = dequeue(&amp;c-&gt;sendq, c);
	if(sg != nil) {
		gp = sg-&gt;g;
		freesg(c, sg);
		ready(gp);
	}
	goto retc;

asyns:
	if(cas-&gt;u.elem != nil)
		c-&gt;elemalg-&gt;copy(c-&gt;elemsize, c-&gt;senddataq-&gt;elem, cas-&gt;u.elem);
	c-&gt;senddataq = c-&gt;senddataq-&gt;link;
	c-&gt;qcount++;
	sg = dequeue(&amp;c-&gt;recvq, c);
	if(sg != nil) {
		gp = sg-&gt;g;
		freesg(c, sg);
		ready(gp);
	}
	goto retc;

gotr:
	// recv path to wakeup the sender (sg)
	if(debug) {
		prints(&#34;gotr: sel=&#34;);
		runtime·printpointer(sel);
		prints(&#34; c=&#34;);
		runtime·printpointer(c);
		prints(&#34; o=&#34;);
		runtime·printint(o);
		prints(&#34;\n&#34;);
	}
	if(cas-&gt;u.elemp != nil)
		c-&gt;elemalg-&gt;copy(c-&gt;elemsize, cas-&gt;u.elemp, sg-&gt;elem);
	gp = sg-&gt;g;
	gp-&gt;param = sg;
	ready(gp);
	goto retc;

rclose:
	if(cas-&gt;u.elemp != nil)
		c-&gt;elemalg-&gt;copy(c-&gt;elemsize, cas-&gt;u.elemp, nil);
	c-&gt;closed |= Rclosed;
	incerr(c);
	goto retc;

gots:
	// send path to wakeup the receiver (sg)
	if(debug) {
		prints(&#34;gots: sel=&#34;);
		runtime·printpointer(sel);
		prints(&#34; c=&#34;);
		runtime·printpointer(c);
		prints(&#34; o=&#34;);
		runtime·printint(o);
		prints(&#34;\n&#34;);
	}
	if(c-&gt;closed &amp; Wclosed)
		goto sclose;
	c-&gt;elemalg-&gt;copy(c-&gt;elemsize, sg-&gt;elem, cas-&gt;u.elem);
	gp = sg-&gt;g;
	gp-&gt;param = sg;
	ready(gp);
	goto retc;

sclose:
	incerr(c);
	goto retc;

retc:
	unlock(&amp;chanlock);

	runtime·setcallerpc(&amp;sel, cas-&gt;pc);
	as = (byte*)&amp;sel + cas-&gt;so;
	freesel(sel);
	*as = true;
}

// closechan(sel *byte);
void
runtime·closechan(Hchan *c)
{
	SudoG *sg;
	G* gp;

	lock(&amp;chanlock);
	incerr(c);
	c-&gt;closed |= Wclosed;

	// release all readers
	for(;;) {
		sg = dequeue(&amp;c-&gt;recvq, c);
		if(sg == nil)
			break;
		gp = sg-&gt;g;
		gp-&gt;param = nil;
		freesg(c, sg);
		ready(gp);
	}

	// release all writers
	for(;;) {
		sg = dequeue(&amp;c-&gt;sendq, c);
		if(sg == nil)
			break;
		gp = sg-&gt;g;
		gp-&gt;param = nil;
		freesg(c, sg);
		ready(gp);
	}

	unlock(&amp;chanlock);
}

void
chanclose(Hchan *c)
{
	runtime·closechan(c);
}

bool
chanclosed(Hchan *c)
{
	return (c-&gt;closed &amp; Rclosed) != 0;
}

int32
chanlen(Hchan *c)
{
	return c-&gt;qcount;
}

int32
chancap(Hchan *c)
{
	return c-&gt;dataqsiz;
}


// closedchan(sel *byte) bool;
void
runtime·closedchan(Hchan *c, bool closed)
{
	closed = chanclosed(c);
	FLUSH(&amp;closed);
}

static SudoG*
dequeue(WaitQ *q, Hchan *c)
{
	SudoG *sgp;

loop:
	sgp = q-&gt;first;
	if(sgp == nil)
		return nil;
	q-&gt;first = sgp-&gt;link;

	// if sgp is stale, ignore it
	if(sgp-&gt;selgen != sgp-&gt;g-&gt;selgen) {
		//prints(&#34;INVALID PSEUDOG POINTER\n&#34;);
		freesg(c, sgp);
		goto loop;
	}

	// invalidate any others
	sgp-&gt;g-&gt;selgen++;
	return sgp;
}

static void
enqueue(WaitQ *q, SudoG *sgp)
{
	sgp-&gt;link = nil;
	if(q-&gt;first == nil) {
		q-&gt;first = sgp;
		q-&gt;last = sgp;
		return;
	}
	q-&gt;last-&gt;link = sgp;
	q-&gt;last = sgp;
}

static SudoG*
allocsg(Hchan *c)
{
	SudoG* sg;

	sg = c-&gt;free;
	if(sg != nil) {
		c-&gt;free = sg-&gt;link;
	} else
		sg = mal(sizeof(*sg) + c-&gt;elemsize - sizeof(sg-&gt;elem));
	sg-&gt;selgen = g-&gt;selgen;
	sg-&gt;g = g;
	sg-&gt;offset = 0;
	sg-&gt;isfree = 0;

	return sg;
}

static void
freesg(Hchan *c, SudoG *sg)
{
	if(sg != nil) {
		if(sg-&gt;isfree)
			throw(&#34;chan.freesg: already free&#34;);
		sg-&gt;isfree = 1;
		sg-&gt;link = c-&gt;free;
		c-&gt;free = sg;
	}
}

static uint32
gcd(uint32 u, uint32 v)
{
	for(;;) {
		if(u &gt; v) {
			if(v == 0)
				return u;
			u = u%v;
			continue;
		}
		if(u == 0)
			return v;
		v = v%u;
	}
}

static uint32
fastrand1(void)
{
	static uint32 x = 0x49f6428aUL;

	x += x;
	if(x &amp; 0x80000000L)
		x ^= 0x88888eefUL;
	return x;
}

static uint32
fastrand2(void)
{
	static uint32 x = 0x49f6428aUL;

	x += x;
	if(x &amp; 0x80000000L)
		x ^= 0xfafd871bUL;
	return x;
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

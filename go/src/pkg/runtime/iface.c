<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/iface.c</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/iface.c</h1>

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

static void
printiface(Iface i)
{
	printf(&#34;(%p,%p)&#34;, i.tab, i.data);
}

static void
printeface(Eface e)
{
	printf(&#34;(%p,%p)&#34;, e.type, e.data);
}

/*
 * layout of Itab known to compilers
 */
struct Itab
{
	InterfaceType*	inter;
	Type*	type;
	Itab*	link;
	int32	bad;
	int32	unused;
	void	(*fun[])(void);
};

static	Itab*	hash[1009];
static	Lock	ifacelock;

static Itab*
itab(InterfaceType *inter, Type *type, int32 canfail)
{
	int32 locked;
	int32 ni;
	Method *t, *et;
	IMethod *i, *ei;
	uint32 ihash, h;
	String *iname;
	Itab *m;
	UncommonType *x;

	if(inter-&gt;mhdr.len == 0)
		throw(&#34;internal error - misuse of itab&#34;);

	// easy case
	x = type-&gt;x;
	if(x == nil) {
		if(canfail)
			return nil;
		iname = inter-&gt;m[0].name;
		goto throw;
	}

	// compiler has provided some good hash codes for us.
	h = inter-&gt;hash;
	h += 17 * type-&gt;hash;
	// TODO(rsc): h += 23 * x-&gt;mhash ?
	h %= nelem(hash);

	// look twice - once without lock, once with.
	// common case will be no lock contention.
	for(locked=0; locked&lt;2; locked++) {
		if(locked)
			lock(&amp;ifacelock);
		for(m=hash[h]; m!=nil; m=m-&gt;link) {
			if(m-&gt;inter == inter &amp;&amp; m-&gt;type == type) {
				if(m-&gt;bad) {
					m = nil;
					if(!canfail) {
						// this can only happen if the conversion
						// was already done once using the , ok form
						// and we have a cached negative result.
						// the cached result doesn&#39;t record which
						// interface function was missing, so jump
						// down to the interface check, which will
						// do more work but give a better error.
						goto search;
					}
				}
				if(locked)
					unlock(&amp;ifacelock);
				return m;
			}
		}
	}

	ni = inter-&gt;mhdr.len;
	m = malloc(sizeof(*m) + ni*sizeof m-&gt;fun[0]);
	m-&gt;inter = inter;
	m-&gt;type = type;

search:
	// both inter and type have method sorted by hash,
	// so can iterate over both in lock step;
	// the loop is O(ni+nt) not O(ni*nt).
	i = inter-&gt;m;
	ei = i + inter-&gt;mhdr.len;
	t = x-&gt;m;
	et = t + x-&gt;mhdr.len;
	for(; i &lt; ei; i++) {
		ihash = i-&gt;hash;
		iname = i-&gt;name;
		for(;; t++) {
			if(t &gt;= et) {
				if(!canfail) {
				throw:
					// didn&#39;t find method
					printf(&#34;%S is not %S: missing method %S\n&#34;,
						*type-&gt;string, *inter-&gt;string, *iname);
					throw(&#34;interface conversion&#34;);
					return nil;	// not reached
				}
				m-&gt;bad = 1;
				goto out;
			}
			if(t-&gt;hash == ihash &amp;&amp; t-&gt;name == iname)
				break;
		}
		if(m)
			m-&gt;fun[i-&gt;perm] = t-&gt;ifn;
	}

out:
	m-&gt;link = hash[h];
	hash[h] = m;
	if(locked)
		unlock(&amp;ifacelock);
	if(m-&gt;bad)
		return nil;
	return m;
}

static void
copyin(Type *t, void *src, void **dst)
{
	int32 wid, alg;
	void *p;

	wid = t-&gt;size;
	alg = t-&gt;alg;

	if(wid &lt;= sizeof(*dst))
		algarray[alg].copy(wid, dst, src);
	else {
		p = mal(wid);
		algarray[alg].copy(wid, p, src);
		*dst = p;
	}
}

static void
copyout(Type *t, void **src, void *dst)
{
	int32 wid, alg;

	wid = t-&gt;size;
	alg = t-&gt;alg;

	if(wid &lt;= sizeof(*src))
		algarray[alg].copy(wid, dst, src);
	else
		algarray[alg].copy(wid, dst, *src);
}

// ifaceT2I(sigi *byte, sigt *byte, elem any) (ret Iface);
#pragma textflag 7
void
runtime·ifaceT2I(InterfaceType *inter, Type *t, ...)
{
	byte *elem;
	Iface *ret;
	int32 wid;

	elem = (byte*)(&amp;t+1);
	wid = t-&gt;size;
	ret = (Iface*)(elem + rnd(wid, Structrnd));
	ret-&gt;tab = itab(inter, t, 0);
	copyin(t, elem, &amp;ret-&gt;data);
}

// ifaceT2E(sigt *byte, elem any) (ret Eface);
#pragma textflag 7
void
runtime·ifaceT2E(Type *t, ...)
{
	byte *elem;
	Eface *ret;
	int32 wid;

	elem = (byte*)(&amp;t+1);
	wid = t-&gt;size;
	ret = (Eface*)(elem + rnd(wid, Structrnd));

	ret-&gt;type = t;
	copyin(t, elem, &amp;ret-&gt;data);
}

// ifaceI2T(sigt *byte, iface any) (ret any);
#pragma textflag 7
void
runtime·ifaceI2T(Type *t, Iface i, ...)
{
	Itab *tab;
	byte *ret;

	ret = (byte*)(&amp;i+1);
	tab = i.tab;
	if(tab == nil) {
		printf(&#34;interface is nil, not %S\n&#34;, *t-&gt;string);
		throw(&#34;interface conversion&#34;);
	}
	if(tab-&gt;type != t) {
		printf(&#34;%S is %S, not %S\n&#34;, *tab-&gt;inter-&gt;string, *tab-&gt;type-&gt;string, *t-&gt;string);
		throw(&#34;interface conversion&#34;);
	}
	copyout(t, &amp;i.data, ret);
}

// ifaceI2T2(sigt *byte, i Iface) (ret any, ok bool);
#pragma textflag 7
void
runtime·ifaceI2T2(Type *t, Iface i, ...)
{
	byte *ret;
	bool *ok;
	int32 wid;

	ret = (byte*)(&amp;i+1);
	wid = t-&gt;size;
	ok = (bool*)(ret+rnd(wid, 1));

	if(i.tab == nil || i.tab-&gt;type != t) {
		*ok = false;
		runtime·memclr(ret, wid);
		return;
	}

	*ok = true;
	copyout(t, &amp;i.data, ret);
}

// ifaceE2T(sigt *byte, e Eface) (ret any);
#pragma textflag 7
void
runtime·ifaceE2T(Type *t, Eface e, ...)
{
	byte *ret;

	ret = (byte*)(&amp;e+1);

	if(e.type != t) {
		if(e.type == nil)
			printf(&#34;interface is nil, not %S\n&#34;, *t-&gt;string);
		else
			printf(&#34;interface is %S, not %S\n&#34;, *e.type-&gt;string, *t-&gt;string);
		throw(&#34;interface conversion&#34;);
	}
	copyout(t, &amp;e.data, ret);
}

// ifaceE2T2(sigt *byte, iface any) (ret any, ok bool);
#pragma textflag 7
void
runtime·ifaceE2T2(Type *t, Eface e, ...)
{
	byte *ret;
	bool *ok;
	int32 wid;

	ret = (byte*)(&amp;e+1);
	wid = t-&gt;size;
	ok = (bool*)(ret+rnd(wid, 1));

	if(t != e.type) {
		*ok = false;
		runtime·memclr(ret, wid);
		return;
	}

	*ok = true;
	copyout(t, &amp;e.data, ret);
}

// ifaceI2E(sigi *byte, iface any) (ret any);
// TODO(rsc): Move to back end, throw away function.
void
runtime·ifaceI2E(Iface i, Eface ret)
{
	Itab *tab;

	ret.data = i.data;
	tab = i.tab;
	if(tab == nil)
		ret.type = nil;
	else
		ret.type = tab-&gt;type;
	FLUSH(&amp;ret);
}

// ifaceI2I(sigi *byte, iface any) (ret any);
// called only for implicit (no type assertion) conversions.
// converting nil is okay.
void
runtime·ifaceI2I(InterfaceType *inter, Iface i, Iface ret)
{
	Itab *tab;

	tab = i.tab;
	if(tab == nil) {
		// If incoming interface is uninitialized (zeroed)
		// make the outgoing interface zeroed as well.
		ret.tab = nil;
		ret.data = nil;
	} else {
		ret = i;
		if(tab-&gt;inter != inter)
			ret.tab = itab(inter, tab-&gt;type, 0);
	}

	FLUSH(&amp;ret);
}

// ifaceI2Ix(sigi *byte, iface any) (ret any);
// called only for explicit conversions (with type assertion).
// converting nil is not okay.
void
runtime·ifaceI2Ix(InterfaceType *inter, Iface i, Iface ret)
{
	Itab *tab;

	tab = i.tab;
	if(tab == nil) {
		// explicit conversions require non-nil interface value.
		printf(&#34;interface is nil, not %S\n&#34;, *inter-&gt;string);
		throw(&#34;interface conversion&#34;);
	} else {
		ret = i;
		if(tab-&gt;inter != inter)
			ret.tab = itab(inter, tab-&gt;type, 0);
	}

	FLUSH(&amp;ret);
}

// ifaceI2I2(sigi *byte, iface any) (ret any, ok bool);
void
runtime·ifaceI2I2(InterfaceType *inter, Iface i, Iface ret, bool ok)
{
	Itab *tab;

	tab = i.tab;
	if(tab == nil) {
		// If incoming interface is nil, the conversion fails.
		ret.tab = nil;
		ret.data = nil;
		ok = false;
	} else {
		ret = i;
		ok = true;
		if(tab-&gt;inter != inter) {
			ret.tab = itab(inter, tab-&gt;type, 1);
			if(ret.tab == nil) {
				ret.data = nil;
				ok = false;
			}
		}
	}

	FLUSH(&amp;ret);
	FLUSH(&amp;ok);
}

// ifaceE2I(sigi *byte, iface any) (ret any);
// Called only for explicit conversions (with type assertion).
void
ifaceE2I(InterfaceType *inter, Eface e, Iface *ret)
{
	Type *t;

	t = e.type;
	if(t == nil) {
		// explicit conversions require non-nil interface value.
		printf(&#34;interface is nil, not %S\n&#34;, *inter-&gt;string);
		throw(&#34;interface conversion&#34;);
	} else {
		ret-&gt;data = e.data;
		ret-&gt;tab = itab(inter, t, 0);
	}
}

// ifaceE2I(sigi *byte, iface any) (ret any);
// Called only for explicit conversions (with type assertion).
void
runtime·ifaceE2I(InterfaceType *inter, Eface e, Iface ret)
{
	ifaceE2I(inter, e, &amp;ret);
}

// ifaceE2I2(sigi *byte, iface any) (ret any, ok bool);
void
runtime·ifaceE2I2(InterfaceType *inter, Eface e, Iface ret, bool ok)
{
	Type *t;

	t = e.type;
	ok = true;
	if(t == nil) {
		// If incoming interface is nil, the conversion fails.
		ret.data = nil;
		ret.tab = nil;
		ok = false;
	} else {
		ret.data = e.data;
		ret.tab = itab(inter, t, 1);
		if(ret.tab == nil) {
			ret.data = nil;
			ok = false;
		}
	}
	FLUSH(&amp;ret);
	FLUSH(&amp;ok);
}

static uintptr
ifacehash1(void *data, Type *t)
{
	int32 alg, wid;

	if(t == nil)
		return 0;

	alg = t-&gt;alg;
	wid = t-&gt;size;
	if(algarray[alg].hash == nohash) {
		// calling nohash will throw too,
		// but we can print a better error.
		printf(&#34;hash of unhashable type %S\n&#34;, *t-&gt;string);
		if(alg == AFAKE)
			throw(&#34;fake interface hash&#34;);
		throw(&#34;interface hash&#34;);
	}
	if(wid &lt;= sizeof(data))
		return algarray[alg].hash(wid, &amp;data);
	return algarray[alg].hash(wid, data);
}

uintptr
ifacehash(Iface a)
{
	if(a.tab == nil)
		return 0;
	return ifacehash1(a.data, a.tab-&gt;type);
}

uintptr
efacehash(Eface a)
{
	return ifacehash1(a.data, a.type);
}

static bool
ifaceeq1(void *data1, void *data2, Type *t)
{
	int32 alg, wid;

	alg = t-&gt;alg;
	wid = t-&gt;size;

	if(algarray[alg].equal == noequal) {
		// calling noequal will throw too,
		// but we can print a better error.
		printf(&#34;comparing uncomparable type %S\n&#34;, *t-&gt;string);
		if(alg == AFAKE)
			throw(&#34;fake interface compare&#34;);
		throw(&#34;interface compare&#34;);
	}

	if(wid &lt;= sizeof(data1))
		return algarray[alg].equal(wid, &amp;data1, &amp;data2);
	return algarray[alg].equal(wid, data1, data2);
}

bool
ifaceeq(Iface i1, Iface i2)
{
	if(i1.tab != i2.tab)
		return false;
	if(i1.tab == nil)
		return true;
	return ifaceeq1(i1.data, i2.data, i1.tab-&gt;type);
}

bool
efaceeq(Eface e1, Eface e2)
{
	if(e1.type != e2.type)
		return false;
	if(e1.type == nil)
		return true;
	return ifaceeq1(e1.data, e2.data, e1.type);
}

// ifaceeq(i1 any, i2 any) (ret bool);
void
runtime·ifaceeq(Iface i1, Iface i2, bool ret)
{
	ret = ifaceeq(i1, i2);
	FLUSH(&amp;ret);
}

// efaceeq(i1 any, i2 any) (ret bool)
void
runtime·efaceeq(Eface e1, Eface e2, bool ret)
{
	ret = efaceeq(e1, e2);
	FLUSH(&amp;ret);
}

// ifacethash(i1 any) (ret uint32);
void
runtime·ifacethash(Iface i1, uint32 ret)
{
	Itab *tab;

	ret = 0;
	tab = i1.tab;
	if(tab != nil)
		ret = tab-&gt;type-&gt;hash;
	FLUSH(&amp;ret);
}

// efacethash(e1 any) (ret uint32)
void
runtime·efacethash(Eface e1, uint32 ret)
{
	Type *t;

	ret = 0;
	t = e1.type;
	if(t != nil)
		ret = t-&gt;hash;
	FLUSH(&amp;ret);
}

void
runtime·printiface(Iface i)
{
	printiface(i);
}

void
runtime·printeface(Eface e)
{
	printeface(e);
}

void
unsafe·Typeof(Eface e, Eface ret)
{
	if(e.type == nil) {
		ret.type = nil;
		ret.data = nil;
	} else
		ret = *(Eface*)e.type;
	FLUSH(&amp;ret);
}

void
unsafe·Reflect(Eface e, Eface rettype, void *retaddr)
{
	uintptr *p;
	uintptr x;

	if(e.type == nil) {
		rettype.type = nil;
		rettype.data = nil;
		retaddr = 0;
	} else {
		rettype = *(Eface*)e.type;
		if(e.type-&gt;size &lt;= sizeof(uintptr)) {
			// Copy data into x ...
			x = 0;
			algarray[e.type-&gt;alg].copy(e.type-&gt;size, &amp;x, &amp;e.data);

			// but then build pointer to x so that Reflect
			// always returns pointer to data.
			p = mallocgc(sizeof(uintptr));
			*p = x;
		} else {
			// Already a pointer, but still make a copy,
			// to preserve value semantics for interface data.
			p = mallocgc(e.type-&gt;size);
			algarray[e.type-&gt;alg].copy(e.type-&gt;size, p, e.data);
		}
		retaddr = p;
	}
	FLUSH(&amp;rettype);
	FLUSH(&amp;retaddr);
}

void
unsafe·Unreflect(Iface typ, void *addr, Eface e)
{
	// Reflect library has reinterpreted typ
	// as its own kind of type structure.
	// We know that the pointer to the original
	// type structure sits before the data pointer.
	e.type = (Type*)((Eface*)typ.data-1);

	// Interface holds either pointer to data
	// or copy of original data.
	if(e.type-&gt;size &lt;= sizeof(uintptr))
		algarray[e.type-&gt;alg].copy(e.type-&gt;size, &amp;e.data, addr);
	else {
		// Easier: already a pointer to data.
		// TODO(rsc): Should this make a copy?
		e.data = addr;
	}

	FLUSH(&amp;e);
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

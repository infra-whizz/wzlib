# Overview

Query parser for Whizz.

#Query syntax

This query parser has similar syntax from
[Sugar Project](https://github.com/sugarsack/sugar):

	[flag]:[trait]:<target>

## Flags

Flags are:

- `x`: revert the expression
- `a`: is reserved and is equal to `*`

## Traits

Trait is a key from the traits and thus turns `target` into a value of
trait. Example:

	os:debian

# Operators: AND, OR

You can join matcher within your query with and and or operators. To
use `not` operator you should set an inverter flag `x` to a matcher,
so your expression with be inverted, such as `and` to `and not ...` as
well as `or` to `or not ...` syntax.

## Joining with AND operator

To join any subsequent matcher into one chain with `and` operator, use `/`
(slash) or `&&` or just ` and ` with spaces around. The following examples
are equally same:

	<first>/<second>
	<first>&&<second>
	<first> and <second>

Note, that `and` syntax requires at leat one space in front and after.

## Joining with OR operator

The `or` operator works between blocks over the entire set of known
machines. To join set of `and`-joined matchers or standalone matchers
with `or` operator, use `//` (double-slash), `||` or just ` or ` with
spaces around. The following examples are equally same: 

	<first>//<second>
	<first>||<second>
	<first> or <second>

Here *"second"* expression will be picking from the same original
source as *"first"*, and then both results will be combined together.

## Invert results with NOT operator

To invert something, use `x` flag:

	:x:<expression>

In this case the result of the *"expression"* will be inverted. For
example, match all systems but foo: 

	whizz 'x::foo' ansible.system.ping

Please note here that flag `x::foo` contains colon `:` twice. In this
case Sugar interprets `x` as a flag. However, if the expression would
be just `x:foo`, then the `x` would be interpreted as a trait key `x`
and target `foo` as a trait value.

# Examples

This example is using all three ways of writing logical operators. All
further examples will use slashes-based syntax.

The following string matches all Debian clients with a hostname that
begins with `webserv`, as well as any machines that have a hostname
which matches the regular expression `web-dc1-srv.*`: 

	whizz 'webserv* / os:debian // web-dc1-srv.*' ansible.system.ping
	whizz 'webserv* && os:debian || web-dc1-srv.*' ansible.system.ping
	whizz 'webserv* and os:debian or web-dc1-srv.*' ansible.system.ping

Spaces between logical operators might be omitted, except and and or ones.

Inversion works through the flag `x`. So in order to exclude a client
hostname (ping all machines, except `web-dc1-srv`):

	whizz x::web-dc1-srv ansible.system.ping


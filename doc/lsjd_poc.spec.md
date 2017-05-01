Libre Science Journal Daemon
============================
A Proof of Concept

This specification is a work in progress

Author: R. Díaz de León Plata <leon@elinter.net>


## Notice

The end goal is to make a full p2p service, this version however is going to
use the ljsd as a centralized HTTP daemon for simplicity.

## Scenarios

Someone wanting to submit an article for review should sing it and send it
to the lsjd server using the lsj CLI tool.

Researchers should browse the lsjd listings in search of new documents
to read and verify.

## Non-goal

- Managing tags
- Sharing keys
- Verifying keys
- Support for web browsers

## Overview

The lsjd will be used to upload, browse and search for articles through
the lsj command line tool.


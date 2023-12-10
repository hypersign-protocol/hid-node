#!/bin/bash

tmux list-sessions
tmux kill-session -t node1
tmux kill-session -t node2
tmux kill-session -t node3


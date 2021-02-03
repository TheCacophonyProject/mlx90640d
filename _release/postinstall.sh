#!/bin/bash
systemctl daemon-reload
systemctl enable mlx90640d.service
systemctl restart mlx90640d.service
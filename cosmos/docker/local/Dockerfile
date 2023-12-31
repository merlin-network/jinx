
# syntax=docker/dockerfile:1
# Copyright (C) 2022, Blackchain Foundation. All rights reserved.
# See the file LICENSE for licensing terms.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

ARG GO_VERSION
ARG GOARCH=arm64
ARG GOOS=darwin
ARG JINXD_VERSION=v0.0.0

#######################################################
###         Stage 1 - Build Smart Contracts         ###
#######################################################

# Use the latest foundry image
FROM jinxd-base:${JINXD_VERSION} as localjinxd

RUN apk add --no-cache bash jq

WORKDIR /

# Copy over the local information
COPY ./cosmos/docker/local/docker-init.sh /docker-init.sh
COPY ./cosmos/docker/local/config /config

# Set the hom directory
ENV HOME /
WORKDIR $HOME

CMD ["bash", "docker-init.sh"]
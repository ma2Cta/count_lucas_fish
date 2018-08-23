FROM centos:latest

RUN yum install -y wget && \
    wget https://download.opensuse.org/repositories/shells:/fish:/release:/2/CentOS_7/x86_64/fish-2.7.1-1.2.x86_64.rpm && \
    yum localinstall -y fish-2.7.1-1.2.x86_64.rpm

CMD /bin/fish

FROM centos

RUN dnf --disablerepo '*' --enablerepo=extras swap centos-linux-repos centos-stream-repos -y \
    && dnf distro-sync -y \
    && yum install python3.8 -y \
    && dnf install @mysql -y \
    && yum install golang -y

CMD ["systemctl", "enable", "--now mysqld"]
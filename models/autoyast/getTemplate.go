package autoyast

func GetTemplate() (temp string) {
	temp = `<?xml version="1.0"?>
<!DOCTYPE profile>
<profile xmlns="http://www.suse.com/1.0/yast2ns" xmlns:config="http://www.suse.com/1.0/configns">
  <add-on>
    <add_on_products config:type="list"/>
  </add-on>
  <bootloader>
    <global>
      <activate>false</activate>
      <append>hvc_iucv=8 TERM=dumb resume=/dev/mapper/example_lunid-part2 cio_ignore=all,!ipldev,!condev crashkernel=164M</append>
      <boot_boot>false</boot_boot>
      <boot_extended>false</boot_extended>
      <boot_mbr>false</boot_mbr>
      <boot_root>false</boot_root>
      <generic_mbr>false</generic_mbr>
      <gfxmode>auto</gfxmode>
      <hiddenmenu>false</hiddenmenu>
      <os_prober>false</os_prober>
      <terminal>console</terminal>
      <timeout config:type="integer">8</timeout>
      <trusted_grub>false</trusted_grub>
      <xen_append>crashkernel=164M</xen_append>
      <xen_kernel_append>crashkernel=164M\&lt;4G</xen_kernel_append>
    </global>
    <loader_type>grub2</loader_type>
  </bootloader>
  <ca_mgm>
    <CAName>YaST_Default_CA</CAName>
    <ca_commonName>YaST Default CA ()</ca_commonName>
    <country>US</country>
    <locality/>
    <organisation/>
    <organisationUnit/>
    <password>ENTER PASSWORD HERE</password>
    <server_email>postmaster@</server_email>
    <state/>
    <takeLocalServerName config:type="boolean">true</takeLocalServerName>
  </ca_mgm>
  <dasd>
    <devices config:type="list"/>
    <format_unformatted config:type="boolean">false</format_unformatted>
  </dasd>
  <deploy_image>
    <image_installation config:type="boolean">false</image_installation>
  </deploy_image>
  <firewall>
    <FW_ALLOW_FW_BROADCAST_DMZ>no</FW_ALLOW_FW_BROADCAST_DMZ>
    <FW_ALLOW_FW_BROADCAST_EXT>no</FW_ALLOW_FW_BROADCAST_EXT>
    <FW_ALLOW_FW_BROADCAST_INT>no</FW_ALLOW_FW_BROADCAST_INT>
    <FW_BOOT_FULL_INIT>no</FW_BOOT_FULL_INIT>
    <FW_CONFIGURATIONS_DMZ>vnc-httpd vnc-server</FW_CONFIGURATIONS_DMZ>
    <FW_CONFIGURATIONS_EXT>vnc-httpd vnc-server</FW_CONFIGURATIONS_EXT>
    <FW_CONFIGURATIONS_INT>vnc-httpd vnc-server</FW_CONFIGURATIONS_INT>
    <FW_DEV_DMZ/>
    <FW_DEV_EXT/>
    <FW_DEV_INT/>
    <FW_FORWARD_ALWAYS_INOUT_DEV/>
    <FW_FORWARD_MASQ/>
    <FW_IGNORE_FW_BROADCAST_DMZ>no</FW_IGNORE_FW_BROADCAST_DMZ>
    <FW_IGNORE_FW_BROADCAST_EXT>yes</FW_IGNORE_FW_BROADCAST_EXT>
    <FW_IGNORE_FW_BROADCAST_INT>no</FW_IGNORE_FW_BROADCAST_INT>
    <FW_IPSEC_TRUST>no</FW_IPSEC_TRUST>
    <FW_LOAD_MODULES>nf_conntrack_netbios_ns</FW_LOAD_MODULES>
    <FW_LOG_ACCEPT_ALL>no</FW_LOG_ACCEPT_ALL>
    <FW_LOG_ACCEPT_CRIT>yes</FW_LOG_ACCEPT_CRIT>
    <FW_LOG_DROP_ALL>no</FW_LOG_DROP_ALL>
    <FW_LOG_DROP_CRIT>yes</FW_LOG_DROP_CRIT>
    <FW_MASQUERADE>no</FW_MASQUERADE>
    <FW_PROTECT_FROM_INT>no</FW_PROTECT_FROM_INT>
    <FW_ROUTE>no</FW_ROUTE>
    <FW_SERVICES_ACCEPT_DMZ/>
    <FW_SERVICES_ACCEPT_EXT/>
    <FW_SERVICES_ACCEPT_INT/>
    <FW_SERVICES_ACCEPT_RELATED_DMZ/>
    <FW_SERVICES_ACCEPT_RELATED_EXT/>
    <FW_SERVICES_ACCEPT_RELATED_INT/>
    <FW_SERVICES_DMZ_IP/>
    <FW_SERVICES_DMZ_RPC/>
    <FW_SERVICES_DMZ_TCP/>
    <FW_SERVICES_DMZ_UDP/>
    <FW_SERVICES_EXT_IP/>
    <FW_SERVICES_EXT_RPC/>
    <FW_SERVICES_EXT_TCP/>
    <FW_SERVICES_EXT_UDP/>
    <FW_SERVICES_INT_IP/>
    <FW_SERVICES_INT_RPC/>
    <FW_SERVICES_INT_TCP/>
    <FW_SERVICES_INT_UDP/>
    <FW_STOP_KEEP_ROUTING_STATE>no</FW_STOP_KEEP_ROUTING_STATE>
    <enable_firewall config:type="boolean">false</enable_firewall>
    <start_firewall config:type="boolean">false</start_firewall>
  </firewall>
  <general>
    <ask-list config:type="list"/>
    <cio_ignore config:type="boolean">true</cio_ignore>
    <mode>
      <confirm config:type="boolean">false</confirm>
    </mode>
    <proposals config:type="list"/>
    <signature-handling>
      <accept_file_without_checksum config:type="boolean">true</accept_file_without_checksum>
      <accept_non_trusted_gpg_key config:type="boolean">true</accept_non_trusted_gpg_key>
      <accept_unknown_gpg_key config:type="boolean">true</accept_unknown_gpg_key>
      <accept_unsigned_file config:type="boolean">true</accept_unsigned_file>
      <accept_verification_failed config:type="boolean">false</accept_verification_failed>
      <import_gpg_key config:type="boolean">true</import_gpg_key>
    </signature-handling>
    <storage>
      <partition_alignment config:type="symbol">align_optimal</partition_alignment>
      <start_multipath config:type="boolean">true</start_multipath>
    </storage>
  </general>
  <groups config:type="list">
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>100</gid>
      <group_password>x</group_password>
      <groupname>users</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>0</gid>
      <group_password>x</group_password>
      <groupname>root</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>2</gid>
      <group_password>x</group_password>
      <groupname>daemon</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>482</gid>
      <group_password>x</group_password>
      <groupname>gdm</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>492</gid>
      <group_password>x</group_password>
      <groupname>systemd-timesync</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>43</gid>
      <group_password>x</group_password>
      <groupname>modem</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>42</gid>
      <group_password>x</group_password>
      <groupname>trusted</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>488</gid>
      <group_password>x</group_password>
      <groupname>scard</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>32</gid>
      <group_password>x</group_password>
      <groupname>public</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>486</gid>
      <group_password>x</group_password>
      <groupname>brlapi</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>491</gid>
      <group_password>x</group_password>
      <groupname>ntp</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>25</gid>
      <group_password>x</group_password>
      <groupname>at</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>59</gid>
      <group_password>x</group_password>
      <groupname>maildrop</groupname>
      <userlist>postfix</userlist>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>13</gid>
      <group_password>x</group_password>
      <groupname>news</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>41</gid>
      <group_password>x</group_password>
      <groupname>xok</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>54</gid>
      <group_password>x</group_password>
      <groupname>lock</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>481</gid>
      <group_password>x</group_password>
      <groupname>adm</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>5</gid>
      <group_password>x</group_password>
      <groupname>tty</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>10</gid>
      <group_password>x</group_password>
      <groupname>wheel</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>20</gid>
      <group_password>x</group_password>
      <groupname>cdrom</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>62</gid>
      <group_password>x</group_password>
      <groupname>man</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>484</gid>
      <group_password>x</group_password>
      <groupname>pulse-access</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>49</gid>
      <group_password>x</group_password>
      <groupname>ftp</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>7</gid>
      <group_password>x</group_password>
      <groupname>lp</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>494</gid>
      <group_password>x</group_password>
      <groupname>nscd</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>483</gid>
      <group_password>x</group_password>
      <groupname>vnc</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>490</gid>
      <group_password>x</group_password>
      <groupname>ts-shell</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>19</gid>
      <group_password>x</group_password>
      <groupname>floppy</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>21</gid>
      <group_password>x</group_password>
      <groupname>console</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>3</gid>
      <group_password>x</group_password>
      <groupname>sys</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>496</gid>
      <group_password>x</group_password>
      <groupname>input</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>51</gid>
      <group_password>x</group_password>
      <groupname>postfix</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>12</gid>
      <group_password>x</group_password>
      <groupname>mail</groupname>
      <userlist>postfix</userlist>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>487</gid>
      <group_password>x</group_password>
      <groupname>rtkit</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>22</gid>
      <group_password>x</group_password>
      <groupname>utmp</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>6</gid>
      <group_password>x</group_password>
      <groupname>disk</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>16</gid>
      <group_password>x</group_password>
      <groupname>dialout</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>498</gid>
      <group_password>x</group_password>
      <groupname>sshd</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>497</gid>
      <group_password>x</group_password>
      <groupname>tape</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>15</gid>
      <group_password>x</group_password>
      <groupname>shadow</groupname>
      <userlist>vnc</userlist>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>14</gid>
      <group_password>x</group_password>
      <groupname>uucp</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>495</gid>
      <group_password>x</group_password>
      <groupname>polkitd</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>8</gid>
      <group_password>x</group_password>
      <groupname>www</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>33</gid>
      <group_password>x</group_password>
      <groupname>video</groupname>
      <userlist>gdm</userlist>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>9</gid>
      <group_password>x</group_password>
      <groupname>kmem</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>485</gid>
      <group_password>x</group_password>
      <groupname>pulse</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>65534</gid>
      <group_password>x</group_password>
      <groupname>nogroup</groupname>
      <userlist>nobody</userlist>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>17</gid>
      <group_password>x</group_password>
      <groupname>audio</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>1</gid>
      <group_password>x</group_password>
      <groupname>bin</groupname>
      <userlist>daemon</userlist>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>71</gid>
      <group_password>x</group_password>
      <groupname>ntadmin</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>499</gid>
      <group_password>x</group_password>
      <groupname>messagebus</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>65533</gid>
      <group_password>x</group_password>
      <groupname>nobody</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>489</gid>
      <group_password>x</group_password>
      <groupname>zkeyadm</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>40</gid>
      <group_password>x</group_password>
      <groupname>games</groupname>
      <userlist/>
    </group>
    <group>
      <encrypted config:type="boolean">true</encrypted>
      <gid>493</gid>
      <group_password>x</group_password>
      <groupname>systemd-journal</groupname>
      <userlist/>
    </group>
  </groups>
  <host>
    <hosts config:type="list">
      <hosts_entry>
        <host_address>127.0.0.1</host_address>
        <names config:type="list">
          <name>localhost</name>
        </names>
      </hosts_entry>
      <hosts_entry>
        <host_address>example_ip</host_address>
        <names config:type="list">
          <name>linux-vuqq</name>
        </names>
      </hosts_entry>
      <hosts_entry>
        <host_address>::1</host_address>
        <names config:type="list">
          <name>localhost ipv6-localhost ipv6-loopback</name>
        </names>
      </hosts_entry>
      <hosts_entry>
        <host_address>fe00::0</host_address>
        <names config:type="list">
          <name>ipv6-localnet</name>
        </names>
      </hosts_entry>
      <hosts_entry>
        <host_address>ff00::0</host_address>
        <names config:type="list">
          <name>ipv6-mcastprefix</name>
        </names>
      </hosts_entry>
      <hosts_entry>
        <host_address>ff02::1</host_address>
        <names config:type="list">
          <name>ipv6-allnodes</name>
        </names>
      </hosts_entry>
      <hosts_entry>
        <host_address>ff02::2</host_address>
        <names config:type="list">
          <name>ipv6-allrouters</name>
        </names>
      </hosts_entry>
      <hosts_entry>
        <host_address>ff02::3</host_address>
        <names config:type="list">
          <name>ipv6-allhosts</name>
        </names>
      </hosts_entry>
    </hosts>
  </host>
  <kdump>
    <add_crash_kernel config:type="boolean">true</add_crash_kernel>
    <crash_kernel>164M</crash_kernel>
    <crash_xen_kernel>164M\&lt;4G</crash_xen_kernel>
    <general>
      <KDUMPTOOL_FLAGS/>
      <KDUMP_COMMANDLINE/>
      <KDUMP_COMMANDLINE_APPEND/>
      <KDUMP_CONTINUE_ON_ERROR>true</KDUMP_CONTINUE_ON_ERROR>
      <KDUMP_COPY_KERNEL>yes</KDUMP_COPY_KERNEL>
      <KDUMP_CPUS/>
      <KDUMP_DUMPFORMAT>lzo</KDUMP_DUMPFORMAT>
      <KDUMP_DUMPLEVEL>31</KDUMP_DUMPLEVEL>
      <KDUMP_FREE_DISK_SIZE>64</KDUMP_FREE_DISK_SIZE>
      <KDUMP_HOST_KEY/>
      <KDUMP_IMMEDIATE_REBOOT>yes</KDUMP_IMMEDIATE_REBOOT>
      <KDUMP_KEEP_OLD_DUMPS>5</KDUMP_KEEP_OLD_DUMPS>
      <KDUMP_KERNELVER/>
      <KDUMP_NETCONFIG>auto</KDUMP_NETCONFIG>
      <KDUMP_NET_TIMEOUT>30</KDUMP_NET_TIMEOUT>
      <KDUMP_NOTIFICATION_CC/>
      <KDUMP_NOTIFICATION_TO/>
      <KDUMP_POSTSCRIPT/>
      <KDUMP_PRESCRIPT/>
      <KDUMP_REQUIRED_PROGRAMS/>
      <KDUMP_SAVEDIR>/var/crash</KDUMP_SAVEDIR>
      <KDUMP_SMTP_PASSWORD/>
      <KDUMP_SMTP_SERVER/>
      <KDUMP_SMTP_USER/>
      <KDUMP_TRANSFER/>
      <KDUMP_VERBOSE>3</KDUMP_VERBOSE>
      <KEXEC_OPTIONS/>
    </general>
  </kdump>
  <language>
    <language>en_US</language>
    <languages/>
  </language>
  <login_settings/>
  <networking>
    <dhcp_options>
      <dhclient_client_id/>
      <dhclient_hostname_option>AUTO</dhclient_hostname_option>
    </dhcp_options>
    <dns>
      <dhcp_hostname config:type="boolean">true</dhcp_hostname>
      <hostname>linux-vuqq</hostname>
      <resolv_conf_policy>auto</resolv_conf_policy>
      <write_hostname config:type="boolean">false</write_hostname>
    </dns>
    <interfaces config:type="list">
      <interface>
        <bootproto>static</bootproto>
        <device>eth0</device>
        <ipaddr>example_ip</ipaddr>
        <netmask>example_netmask</netmask>
        <prefixlen>example_prefixlen</prefixlen>
        <startmode>auto</startmode>
      </interface>
      <interface>
        <bootproto>static</bootproto>
        <device>lo</device>
        <firewall>no</firewall>
        <ipaddr>127.0.0.1</ipaddr>
        <netmask>255.0.0.0</netmask>
        <network>127.0.0.0</network>
        <prefixlen>8</prefixlen>
        <startmode>nfsroot</startmode>
        <usercontrol>no</usercontrol>
      </interface>
    </interfaces>
    <ipv6 config:type="boolean">true</ipv6>
    <keep_install_network config:type="boolean">true</keep_install_network>
    <managed config:type="boolean">false</managed>
    <net-udev config:type="list">
      <rule>
        <name>eth0</name>
        <rule>KERNELS</rule>
        <value>0.0.3000</value>
      </rule>
    </net-udev>
    <routing>
      <ipv4_forward config:type="boolean">false</ipv4_forward>
      <ipv6_forward config:type="boolean">false</ipv6_forward>
      <routes config:type="list">
        <route>
          <destination>default</destination>
          <device>eth0</device>
          <gateway>example_gateway</gateway>
          <netmask>-</netmask>
        </route>
      </routes>
    </routing>
    <s390-devices config:type="list">
      <listentry>
        <chanids>0.0.3000 0.0.3001 0.0.3002</chanids>
        <layer2 config:type="boolean">true</layer2>
        <portname>no portname required</portname>
        <type>qeth</type>
      </listentry>
      <listentry>
        <chanids>   </chanids>
        <type/>
      </listentry>
    </s390-devices>
  </networking>
  <nis>
    <netconfig_policy>auto</netconfig_policy>
    <nis_broadcast config:type="boolean">false</nis_broadcast>
    <nis_broken_server config:type="boolean">false</nis_broken_server>
    <nis_domain/>
    <nis_local_only config:type="boolean">false</nis_local_only>
    <nis_options/>
    <nis_other_domains config:type="list"/>
    <nis_servers config:type="list"/>
    <slp_domain/>
    <start_autofs config:type="boolean">false</start_autofs>
    <start_nis config:type="boolean">false</start_nis>
  </nis>
  <ntp-client>
    <ntp_policy>auto</ntp_policy>
    <peers config:type="list">
      <peer>
        <address>/var/lib/ntp/drift/ntp.drift</address>
        <comment># path for drift file</comment>
        <options/>
        <type>driftfile</type>
      </peer>
      <peer>
        <address>/var/log/ntp</address>
        <comment># alternate log file</comment>
        <options/>
        <type>logfile</type>
      </peer>
      <peer>
        <address>/etc/ntp.keys</address>
        <comment># path for keys file</comment>
        <options/>
        <type>keys</type>
      </peer>
      <peer>
        <address>1</address>
        <comment># define trusted keys</comment>
        <options/>
        <type>trustedkey</type>
      </peer>
      <peer>
        <address>1</address>
        <comment># key (7) for accessing server variables</comment>
        <options/>
        <type>requestkey</type>
      </peer>
      <peer>
        <address>1</address>
        <comment># key (6) for accessing server variables</comment>
        <options/>
        <type>controlkey</type>
      </peer>
    </peers>
    <restricts config:type="list">
      <restrict>
        <comment/>
        <mask/>
        <options>ipv6 notrap nomodify nopeer noquery</options>
        <target>default</target>
      </restrict>
      <restrict>
        <comment/>
        <mask/>
        <options/>
        <target>127.0.0.1</target>
      </restrict>
      <restrict>
        <comment/>
        <mask/>
        <options/>
        <target>::1</target>
      </restrict>
    </restricts>
    <start_at_boot config:type="boolean">false</start_at_boot>
    <start_in_chroot config:type="boolean">false</start_in_chroot>
    <sync_interval config:type="integer">5</sync_interval>
    <synchronize_time config:type="boolean">false</synchronize_time>
  </ntp-client>
  <partitioning config:type="list">
    <drive>
      <device>/dev/mapper/example_lunid</device>
      <disklabel>msdos</disklabel>
      <enable_snapshots config:type="boolean">false</enable_snapshots>
      <initialize config:type="boolean">true</initialize>
      <partitions config:type="list">
        <partition>
          <create config:type="boolean">true</create>
          <crypt_fs config:type="boolean">false</crypt_fs>
          <filesystem config:type="symbol">ext4</filesystem>
          <format config:type="boolean">true</format>
          <fstopt>acl,user_xattr</fstopt>
          <loop_fs config:type="boolean">false</loop_fs>
          <mount>/boot/zipl</mount>
          <mountby config:type="symbol">device</mountby>
          <partition_id config:type="integer">131</partition_id>
          <partition_nr config:type="integer">1</partition_nr>
          <partition_type>primary</partition_type>
          <resize config:type="boolean">false</resize>
          <size>1068662272</size>
        </partition>
        <partition>
          <create config:type="boolean">true</create>
          <crypt_fs config:type="boolean">false</crypt_fs>
          <filesystem config:type="symbol">swap</filesystem>
          <format config:type="boolean">true</format>
          <fstopt>defaults</fstopt>
          <loop_fs config:type="boolean">false</loop_fs>
          <mount>swap</mount>
          <mountby config:type="symbol">device</mountby>
          <partition_id config:type="integer">130</partition_id>
          <partition_nr config:type="integer">2</partition_nr>
          <partition_type>primary</partition_type>
          <resize config:type="boolean">false</resize>
          <size>26838466048</size>
        </partition>
        <partition>
          <create config:type="boolean">true</create>
          <crypt_fs config:type="boolean">false</crypt_fs>
          <format config:type="boolean">false</format>
          <loop_fs config:type="boolean">false</loop_fs>
          <lvm_group>rootvg</lvm_group>
          <mountby config:type="symbol">device</mountby>
          <partition_id config:type="integer">142</partition_id>
          <partition_nr config:type="integer">3</partition_nr>
          <partition_type>primary</partition_type>
          <resize config:type="boolean">false</resize>
          <size>186815512064</size>
        </partition>
      </partitions>
      <pesize/>
      <type config:type="symbol">CT_DISK</type>
      <use>all</use>
    </drive>
    <drive>
      <device>/dev/rootvg</device>
      <disklabel>msdos</disklabel>
      <enable_snapshots config:type="boolean">false</enable_snapshots>
      <initialize config:type="boolean">true</initialize>
      <partitions config:type="list">
        <partition>
          <create config:type="boolean">true</create>
          <crypt_fs config:type="boolean">false</crypt_fs>
          <filesystem config:type="symbol">ext4</filesystem>
          <format config:type="boolean">true</format>
          <fstopt>acl,user_xattr</fstopt>
          <loop_fs config:type="boolean">false</loop_fs>
          <lv_name>lv_informix</lv_name>
          <mount>/home/informix</mount>
          <mountby config:type="symbol">device</mountby>
          <partition_nr config:type="integer">1</partition_nr>
          <resize config:type="boolean">false</resize>
          <size>21474836480</size>
        </partition>
        <partition>
          <create config:type="boolean">true</create>
          <crypt_fs config:type="boolean">false</crypt_fs>
          <filesystem config:type="symbol">ext4</filesystem>
          <format config:type="boolean">true</format>
          <fstopt>acl,user_xattr</fstopt>
          <loop_fs config:type="boolean">false</loop_fs>
          <lv_name>lv_informix_tmp</lv_name>
          <mount>/home/informix/tmp</mount>
          <mountby config:type="symbol">device</mountby>
          <partition_nr config:type="integer">1</partition_nr>
          <resize config:type="boolean">false</resize>
          <size>21474836480</size>
        </partition>
        <partition>
          <create config:type="boolean">true</create>
          <crypt_fs config:type="boolean">false</crypt_fs>
          <filesystem config:type="symbol">ext4</filesystem>
          <format config:type="boolean">true</format>
          <fstopt>acl,user_xattr</fstopt>
          <loop_fs config:type="boolean">false</loop_fs>
          <lv_name>lv_piccmon</lv_name>
          <mount>/piccmon</mount>
          <mountby config:type="symbol">device</mountby>
          <partition_nr config:type="integer">1</partition_nr>
          <resize config:type="boolean">false</resize>
          <size>3221225472</size>
        </partition>
        <partition>
          <create config:type="boolean">true</create>
          <crypt_fs config:type="boolean">false</crypt_fs>
          <filesystem config:type="symbol">ext4</filesystem>
          <format config:type="boolean">true</format>
          <fstopt>acl,user_xattr</fstopt>
          <loop_fs config:type="boolean">false</loop_fs>
          <lv_name>lv_root</lv_name>
          <mount>/</mount>
          <mountby config:type="symbol">device</mountby>
          <partition_nr config:type="integer">1</partition_nr>
          <resize config:type="boolean">false</resize>
          <size>129905983488</size>
        </partition>
        <partition>
          <create config:type="boolean">true</create>
          <crypt_fs config:type="boolean">false</crypt_fs>
          <filesystem config:type="symbol">ext4</filesystem>
          <format config:type="boolean">true</format>
          <fstopt>acl,user_xattr</fstopt>
          <loop_fs config:type="boolean">false</loop_fs>
          <lv_name>lv_tsm</lv_name>
          <mount>/opt/tivoli/tsm</mount>
          <mountby config:type="symbol">device</mountby>
          <partition_nr config:type="integer">1</partition_nr>
          <resize config:type="boolean">false</resize>
          <size>10737418240</size>
        </partition>
      </partitions>
      <pesize>4M</pesize>
      <type config:type="symbol">CT_LVM</type>
      <use>all</use>
    </drive>
  </partitioning>
  <printer>
    <client_conf_content>
      <file_contents><![CDATA[# CUPS client configuration file (optional).

# You may use /etc/cups/client.conf (system wide)
# or ~/.cups/client.conf (per user).
# For more information see "man 5 client.conf".

# The ServerName directive specifies the remote server
# that is to be used for all client operations. That is, it
# redirects all client requests directly to that remote server
# so that a local running cupsd is not used in this case.
# The default is to use the local server ("localhost") or domain socket.
# Only one ServerName directive may appear.
# If multiple names are present, only the last one is used.
# The default port number is 631 but can be overridden by adding
# a colon followed by the desired port number.
# The default IPP version is 2.0 but can be overridden by adding
# a slash followed by version=V where V is 1.0 or 1.1 or 2.0 or 2.1 or 2.2.
# IPP version 2.0 does do not work with CUPS 1.3 or older servers.
# If an CUPS 1.3 or older server is used, its older IPP version
# must be specified as .../version=1.1 or .../version=1.0.

# Examples:
# ServerName sever.example.com
# ServerName 192.0.2.10
# ServerName sever.example.com:8631
# ServerName older.server.example.com/version=1.1
# ServerName older.server.example.com:8631/version=1.1

]]></file_contents>
    </client_conf_content>
    <cupsd_conf_content>
      <file_contents><![CDATA[#
# "$Id: cupsd.conf.in 11025 2013-06-07 01:00:33Z msweet $"
#
# Configuration file for the CUPS scheduler.  See "man cupsd.conf" for a
# complete description of this file.
#

# Log general information in error_log - change "warn" to "debug"
# for troubleshooting...
LogLevel warn

# Only listen for connections from the local machine.
Listen localhost:631
Listen /run/cups/cups.sock

# Show shared printers on the local network.
Browsing On
BrowseLocalProtocols dnssd

# Default authentication type, when authentication is required...
DefaultAuthType Basic

# Web interface setting...
WebInterface Yes

# Restrict access to the server...
<Location />
  Order allow,deny
</Location>

# Restrict access to the admin pages...
<Location /admin>
  Order allow,deny
</Location>

# Restrict access to configuration files...
<Location /admin/conf>
  AuthType Default
  Require user @SYSTEM
  Order allow,deny
</Location>

# Set the default printer/job policies...
<Policy default>
  # Job/subscription privacy...
  JobPrivateAccess default
  JobPrivateValues default
  SubscriptionPrivateAccess default
  SubscriptionPrivateValues default

  # Job-related operations must be done by the owner or an administrator...
  <Limit Create-Job Print-Job Print-URI Validate-Job>
    Order deny,allow
  </Limit>

  <Limit Send-Document Send-URI Hold-Job Release-Job Restart-Job Purge-Jobs Set-Job-Attributes Create-Job-Subscription Renew-Subscription Cancel-Subscription Get-Notifications Reprocess-Job Cancel-Current-Job Suspend-Current-Job Resume-Job Cancel-My-Jobs Close-Job CUPS-Move-Job CUPS-Get-Document>
    Require user @OWNER @SYSTEM
    Order deny,allow
  </Limit>

  # All administration operations require an administrator to authenticate...
  <Limit CUPS-Add-Modify-Printer CUPS-Delete-Printer CUPS-Add-Modify-Class CUPS-Delete-Class CUPS-Set-Default CUPS-Get-Devices>
    AuthType Default
    Require user @SYSTEM
    Order deny,allow
  </Limit>

  # All printer operations require a printer operator to authenticate...
  <Limit Pause-Printer Resume-Printer Enable-Printer Disable-Printer Pause-Printer-After-Current-Job Hold-New-Jobs Release-Held-New-Jobs Deactivate-Printer Activate-Printer Restart-Printer Shutdown-Printer Startup-Printer Promote-Job Schedule-Job-After Cancel-Jobs CUPS-Accept-Jobs CUPS-Reject-Jobs>
    AuthType Default
    Require user @SYSTEM
    Order deny,allow
  </Limit>

  # Only the owner or an administrator can cancel or authenticate a job...
  <Limit Cancel-Job CUPS-Authenticate-Job>
    Require user @OWNER @SYSTEM
    Order deny,allow
  </Limit>

  <Limit All>
    Order deny,allow
  </Limit>
</Policy>

# Set the authenticated printer/job policies...
<Policy authenticated>
  # Job/subscription privacy...
  JobPrivateAccess default
  JobPrivateValues default
  SubscriptionPrivateAccess default
  SubscriptionPrivateValues default

  # Job-related operations must be done by the owner or an administrator...
  <Limit Create-Job Print-Job Print-URI Validate-Job>
    AuthType Default
    Order deny,allow
  </Limit>

  <Limit Send-Document Send-URI Hold-Job Release-Job Restart-Job Purge-Jobs Set-Job-Attributes Create-Job-Subscription Renew-Subscription Cancel-Subscription Get-Notifications Reprocess-Job Cancel-Current-Job Suspend-Current-Job Resume-Job Cancel-My-Jobs Close-Job CUPS-Move-Job CUPS-Get-Document>
    AuthType Default
    Require user @OWNER @SYSTEM
    Order deny,allow
  </Limit>

  # All administration operations require an administrator to authenticate...
  <Limit CUPS-Add-Modify-Printer CUPS-Delete-Printer CUPS-Add-Modify-Class CUPS-Delete-Class CUPS-Set-Default>
    AuthType Default
    Require user @SYSTEM
    Order deny,allow
  </Limit>

  # All printer operations require a printer operator to authenticate...
  <Limit Pause-Printer Resume-Printer Enable-Printer Disable-Printer Pause-Printer-After-Current-Job Hold-New-Jobs Release-Held-New-Jobs Deactivate-Printer Activate-Printer Restart-Printer Shutdown-Printer Startup-Printer Promote-Job Schedule-Job-After Cancel-Jobs CUPS-Accept-Jobs CUPS-Reject-Jobs>
    AuthType Default
    Require user @SYSTEM
    Order deny,allow
  </Limit>

  # Only the owner or an administrator can cancel or authenticate a job...
  <Limit Cancel-Job CUPS-Authenticate-Job>
    AuthType Default
    Require user @OWNER @SYSTEM
    Order deny,allow
  </Limit>

  <Limit All>
    Order deny,allow
  </Limit>
</Policy>

# The policy below is added by SUSE during build of our cups package.
# The policy 'allowallforanybody' is totally open and insecure and therefore
# it can only be used within an internal network where only trused users exist
# and where the cupsd is not accessible at all from any external host, see
# http://en.opensuse.org/SDB:CUPS_and_SANE_Firewall_settings
# Have in mind that any user who is allowed to do printer admin tasks
# can change the print queues as he likes - e.g. send copies of confidental
# print jobs from an internal network to any external destination, see
# http://en.opensuse.org/SDB:CUPS_in_a_Nutshell
# For documentation regarding 'Managing Operation Policies' see
# http://www.cups.org/documentation.php/doc-1.7/policies.html
<Policy allowallforanybody>
  # Allow anybody to access job's private values:
  JobPrivateAccess all
  # Make none of the job values to be private:
  JobPrivateValues none
  # Allow anybody to access subscription's private values:
  SubscriptionPrivateAccess all
  # Make none of the subscription values to be private:
  SubscriptionPrivateValues none
  # Allow anybody to do all IPP operations:
  # Currently the IPP operations Validate-Job Cancel-Jobs Cancel-My-Jobs Close-Job CUPS-Get-Document
  # must be additionally exlicitly specified because those IPP operations are not included
  # in the "All" wildcard value - otherwise cupsd prints error messages of the form
  # "No limit for Validate-Job defined in policy allowallforanybody and no suitable template found."
  <Limit All Validate-Job Cancel-Jobs Cancel-My-Jobs Close-Job CUPS-Get-Document>
    Order deny,allow
    Allow from all
  </Limit>
</Policy>
# Explicitly set the CUPS 'default' policy to be used by default:
DefaultPolicy default

#
# End of "$Id: cupsd.conf.in 11025 2013-06-07 01:00:33Z msweet $".
#
]]></file_contents>
    </cupsd_conf_content>
  </printer>
  <proxy>
    <enabled config:type="boolean">false</enabled>
    <ftp_proxy/>
    <http_proxy/>
    <https_proxy/>
    <no_proxy>localhost, 127.0.0.1</no_proxy>
    <proxy_password/>
    <proxy_user/>
  </proxy>
  <report>
    <errors>
      <log config:type="boolean">flase</log>
      <show config:type="boolean">flase</show>
      <timeout config:type="integer">0</timeout>
    </errors>
    <messages>
      <log config:type="boolean">flase</log>
      <show config:type="boolean">flase</show>
      <timeout config:type="integer">0</timeout>
    </messages>
    <warnings>
      <log config:type="boolean">flase</log>
      <show config:type="boolean">flase</show>
      <timeout config:type="integer">0</timeout>
    </warnings>
    <yesno_messages>
      <log config:type="boolean">true</log>
      <show config:type="boolean">true</show>
      <timeout config:type="integer">0</timeout>
    </yesno_messages>
  </report>
  <services-manager>
    <default_target>graphical</default_target>
    <services>
      <disable config:type="list"/>
      <enable config:type="list">
        <service>cio_ignore</service>
        <service>cpi</service>
        <service>cron</service>
        <service>display-manager</service>
        <service>haveged</service>
        <service>iscsi</service>
        <service>kdump-early</service>
        <service>kdump</service>
        <service>multipathd</service>
        <service>nscd</service>
        <service>postfix</service>
        <service>purge-kernels</service>
        <service>rollback</service>
        <service>rsyslog</service>
        <service>smartd</service>
        <service>sshd</service>
        <service>wicked</service>
        <service>wickedd-auto4</service>
        <service>wickedd-dhcp4</service>
        <service>wickedd-dhcp6</service>
        <service>wickedd-nanny</service>
        <service>xinetd</service>
        <service>YaST2-Firstboot</service>
        <service>YaST2-Second-Stage</service>
        <service>getty@tty1</service>
      </enable>
    </services>
  </services-manager>
  <software>
    <image/>
    <install_recommended config:type="boolean">true</install_recommended>
    <instsource/>
    <packages config:type="list">
      <package>xorg-x11-fonts</package>
      <package>xorg-x11-Xvnc</package>
      <package>xorg-x11</package>
      <package>xinetd</package>
      <package>sles-release</package>
      <package>rzsz</package>
      <package>openssh</package>
      <package>lvm2</package>
      <package>kexec-tools</package>
      <package>kdump</package>
      <package>icewm</package>
      <package>grub2</package>
      <package>glibc</package>
      <package>e2fsprogs</package>
    </packages>
    <patterns config:type="list">
      <pattern>32bit</pattern>
      <pattern>Minimal</pattern>
      <pattern>apparmor</pattern>
      <pattern>base</pattern>
      <pattern>documentation</pattern>
      <pattern>gnome-basic</pattern>
      <pattern>sles-Minimal-32bit</pattern>
      <pattern>sles-apparmor-32bit</pattern>
      <pattern>sles-base-32bit</pattern>
      <pattern>sles-documentation-32bit</pattern>
      <pattern>sles-x11-32bit</pattern>
      <pattern>x11</pattern>
      <pattern>yast2</pattern>
    </patterns>
  </software>
  <ssh_import>
    <copy_config config:type="boolean">false</copy_config>
    <import config:type="boolean">false</import>
  </ssh_import>
  <timezone>
    <hwclock>UTC</hwclock>
    <timezone>Asia/Shanghai</timezone>
  </timezone>
  <user_defaults>
    <expire/>
    <group>100</group>
    <groups/>
    <home>/home</home>
    <inactive>-1</inactive>
    <no_groups config:type="boolean">true</no_groups>
    <shell>/bin/bash</shell>
    <skel>/etc/skel</skel>
    <umask>022</umask>
  </user_defaults>
  <users config:type="list">
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>systemd Time Synchronization</fullname>
      <gid>492</gid>
      <home>/</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>492</uid>
      <user_password>!!</user_password>
      <username>systemd-timesync</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Gnome Display Manager daemon</fullname>
      <gid>482</gid>
      <home>/var/lib/gdm</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>484</uid>
      <user_password>!</user_password>
      <username>gdm</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>NFS statd daemon</fullname>
      <gid>65534</gid>
      <home>/var/lib/nfs</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>490</uid>
      <user_password>!</user_password>
      <username>statd</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Daemon</fullname>
      <gid>2</gid>
      <home>/sbin</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>2</uid>
      <user_password>*</user_password>
      <username>daemon</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>root</fullname>
      <gid>0</gid>
      <home>/root</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>0</uid>
      <user_password>$6$9Bp4hAOe9fM0$fYQM4H6uv6OZbAQFPDi.jld0/E8jfgzqP2YAjbfeuWKXM91Lni1i1pl3KWw48eo37ru9cfOcd6u2/i7aOZvYV1</user_password>
      <username>root</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>user for VNC</fullname>
      <gid>483</gid>
      <home>/var/lib/empty</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>485</uid>
      <user_password>!</user_password>
      <username>vnc</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>WWW daemon apache</fullname>
      <gid>8</gid>
      <home>/var/lib/wwwrun</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>30</uid>
      <user_password>*</user_password>
      <username>wwwrun</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>User for GeoClue D-Bus service</fullname>
      <gid>65534</gid>
      <home>/var/lib/srvGeoClue</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>486</uid>
      <user_password>!</user_password>
      <username>srvGeoClue</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>SSH daemon</fullname>
      <gid>498</gid>
      <home>/var/lib/sshd</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>498</uid>
      <user_password>!</user_password>
      <username>sshd</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Secure FTP User</fullname>
      <gid>65534</gid>
      <home>/var/lib/empty</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>491</uid>
      <user_password>!</user_password>
      <username>ftpsecure</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Smart Card Reader</fullname>
      <gid>488</gid>
      <home>/var/run/pcscd</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/usr/sbin/nologin</shell>
      <uid>489</uid>
      <user_password>!</user_password>
      <username>scard</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>openslp daemon</fullname>
      <gid>2</gid>
      <home>/var/lib/empty</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>494</uid>
      <user_password>!</user_password>
      <username>openslp</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Mailer daemon</fullname>
      <gid>12</gid>
      <home>/var/spool/clientmqueue</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>8</uid>
      <user_password>*</user_password>
      <username>mail</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>RealtimeKit</fullname>
      <gid>487</gid>
      <home>/proc</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>488</uid>
      <user_password>!</user_password>
      <username>rtkit</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Postfix Daemon</fullname>
      <gid>51</gid>
      <home>/var/spool/postfix</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>51</uid>
      <user_password>!</user_password>
      <username>postfix</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>User for polkitd</fullname>
      <gid>495</gid>
      <home>/var/lib/polkit</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>497</uid>
      <user_password>!</user_password>
      <username>polkitd</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Unix-to-Unix CoPy system</fullname>
      <gid>14</gid>
      <home>/etc/uucp</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>10</uid>
      <user_password>*</user_password>
      <username>uucp</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>News system</fullname>
      <gid>13</gid>
      <home>/etc/news</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>9</uid>
      <user_password>*</user_password>
      <username>news</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>user for rpcbind</fullname>
      <gid>65534</gid>
      <home>/var/lib/empty</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>495</uid>
      <user_password>!</user_password>
      <username>rpc</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Batch jobs daemon</fullname>
      <gid>25</gid>
      <home>/var/spool/atjobs</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>25</uid>
      <user_password>!</user_password>
      <username>at</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>NTP daemon</fullname>
      <gid>491</gid>
      <home>/var/lib/ntp</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>74</uid>
      <user_password>!</user_password>
      <username>ntp</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Games account</fullname>
      <gid>100</gid>
      <home>/var/games</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>12</uid>
      <user_password>*</user_password>
      <username>games</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>User for nscd</fullname>
      <gid>494</gid>
      <home>/run/nscd</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>496</uid>
      <user_password>!</user_password>
      <username>nscd</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>User for D-Bus</fullname>
      <gid>499</gid>
      <home>/var/run/dbus</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/false</shell>
      <uid>499</uid>
      <user_password>!</user_password>
      <username>messagebus</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Printing daemon</fullname>
      <gid>7</gid>
      <home>/var/spool/lpd</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>4</uid>
      <user_password>*</user_password>
      <username>lp</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>nobody</fullname>
      <gid>65533</gid>
      <home>/var/lib/nobody</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>65534</uid>
      <user_password>*</user_password>
      <username>nobody</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>FTP account</fullname>
      <gid>49</gid>
      <home>/srv/ftp</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>40</uid>
      <user_password>*</user_password>
      <username>ftp</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>bin</fullname>
      <gid>1</gid>
      <home>/bin</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>1</uid>
      <user_password>*</user_password>
      <username>bin</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>PulseAudio daemon</fullname>
      <gid>485</gid>
      <home>/var/lib/pulseaudio</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/sbin/nologin</shell>
      <uid>487</uid>
      <user_password>!</user_password>
      <username>pulse</username>
    </user>
    <user>
      <encrypted config:type="boolean">true</encrypted>
      <fullname>Manual pages viewer</fullname>
      <gid>62</gid>
      <home>/var/cache/man</home>
      <password_settings>
        <expire/>
        <flag/>
        <inact/>
        <max/>
        <min/>
        <warn/>
      </password_settings>
      <shell>/bin/bash</shell>
      <uid>13</uid>
      <user_password>*</user_password>
      <username>man</username>
    </user>
  </users>
  <zfcp>
    <devices config:type="list">
      <listentry>
        <controller_id>0.0.0002</controller_id>
        <fcp_lun>0x0000000000000000</fcp_lun>
        <wwpn>0x500507605e819ef1</wwpn>
      </listentry>
      <listentry>
        <controller_id>0.0.0001</controller_id>
        <fcp_lun>0x0000000000000000</fcp_lun>
        <wwpn>0x500507605e819ef2</wwpn>
      </listentry>
      <listentry>
        <controller_id>0.0.0004</controller_id>
        <fcp_lun>0x0000000000000000</fcp_lun>
        <wwpn>0x500507605e819ec1</wwpn>
      </listentry>
      <listentry>
        <controller_id>0.0.0003</controller_id>
        <fcp_lun>0x0000000000000000</fcp_lun>
        <wwpn>0x500507605e819ec2</wwpn>
      </listentry>
      <listentry>
        <controller_id>0.0.0002</controller_id>
        <fcp_lun>0x0000000000000000</fcp_lun>
        <wwpn>0x50050760448ab202</wwpn>
      </listentry>
      <listentry>
        <controller_id>0.0.0004</controller_id>
        <fcp_lun>0x0000000000000000</fcp_lun>
        <wwpn>0x50050760448ab202</wwpn>
      </listentry>
      <listentry>
        <controller_id>0.0.0001</controller_id>
        <fcp_lun>0x0000000000000000</fcp_lun>
        <wwpn>0x50050760448ab201</wwpn>
      </listentry>
      <listentry>
        <controller_id>0.0.0003</controller_id>
        <fcp_lun>0x0000000000000000</fcp_lun>
        <wwpn>0x50050760448ab201</wwpn>
      </listentry>
    </devices>
  </zfcp>
</profile>`
	return
}

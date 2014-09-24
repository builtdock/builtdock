:title: Customizing router
:description: Learn how to tune custom Deis settings.

.. _router_settings:

Customizing router
=========================
The following settings are tunable for the :ref:`router` component.

Dependencies
------------
Requires: :ref:`builder <builder_settings>`, :ref:`controller <controller_settings>`

Required by: none

Considerations: none

Settings set by router
--------------------------
The following etcd keys are set by the router component, typically in its /bin/boot script.

===========================              =================================================================================
setting                                  description
===========================              =================================================================================
/builtdock/router/$HOST/host                  IP address of the host running this router (there can be multiple routers)
/builtdock/router/$HOST/port                  port used by this router service (there can be multiple routers) (default: 80)
===========================              =================================================================================

Settings used by router
---------------------------
The following etcd keys are used by the router component.

====================================      =============================================================================================================================================================================================
setting                                   description
====================================      =============================================================================================================================================================================================
/builtdock/domains/*                           domain configuration for applications (set by controller)
/builtdock/services/*                          application configuration (set by application unit files)
/builtdock/builder/host                        host of the builder component (set by builder)
/builtdock/builder/port                        port of the builder component (set by builder)
/builtdock/controller/host                     host of the controller component (set by controller)
/builtdock/controller/port                     port of the controller component (set by controller)
/builtdock/router/bodySize                     nginx body size setting (default: 1m)
/builtdock/router/builder/timeout/connect      proxy_connect_timeout for deis-builder (default: 10000). Unit in miliseconds
/builtdock/router/builder/timeout/read         proxy_read_timeout for deis-builder (default: 1200000). Unit in miliseconds
/builtdock/router/builder/timeout/send         proxy_send_timeout for deis-builder (default: 1200000). Unit in miliseconds
/builtdock/router/builder/timeout/tcp          timeout for deis-builder (default: 1200000). Unit in miliseconds
/builtdock/router/controller/timeout/connect   proxy_connect_timeout for deis-controller (default: 10m)
/builtdock/router/controller/timeout/read      proxy_read_timeout for deis-controller (default: 20m)
/builtdock/router/controller/timeout/send      proxy_send_timeout for deis-controller (default: 20m)
/builtdock/router/gzip                         nginx gzip setting (default: on)
/builtdock/router/gzipHttpVersion              nginx gzipHttpVersion setting (default: 1.0)
/builtdock/router/gzipCompLevel                nginx gzipCompLevel setting (default: 2)
/builtdock/router/gzipProxied                  nginx gzipProxied setting (default: any)
/builtdock/router/gzipVary                     nginx gzipVary setting (default: on)
/builtdock/router/gzipDisable                  nginx gzipDisable setting (default: "msie6")
/builtdock/router/gzipTypes                    nginx gzipTypes setting (default: "application/x-javascript, application/xhtml+xml, application/xml, application/xml+rss, application/json, text/css, text/javascript, text/plain, text/xml")
====================================      =============================================================================================================================================================================================

Using a custom router image
---------------------------
You can use a custom Docker image for the router component instead of the image
supplied with Deis:

.. code-block:: console

    $ etcdctl set /builtdock/router/image myaccount/myimage:latest

This will pull the image from the public Docker registry. You can also pull from a private
registry:

.. code-block:: console

    $ etcdctl set /builtdock/router/image registry.mydomain.org:5000/myaccount/myimage:latest

Be sure that your custom image functions in the same way as the `stock router image`_ shipped with
Deis. Specifically, ensure that it sets and reads appropriate etcd keys.

.. _`stock router image`: https://github.com/builtdock/builtdock/tree/master/router

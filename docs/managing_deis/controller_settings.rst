:title: Customizing controller
:description: Learn how to tune custom Deis settings.

.. _controller_settings:

Customizing controller
=========================
The following settings are tunable for the :ref:`controller` component.

Dependencies
------------
Requires: :ref:`controller <controller_settings>`, :ref:`cache <cache_settings>`, :ref:`database <database_settings>`, :ref:`registry <registry_settings>`

Required by: :ref:`router <router_settings>`

Considerations: must live on the same host as both builder and logger (see `#985`_)

Settings set by controller
--------------------------
The following etcd keys are set by the controller component, typically in its /bin/boot script.

===========================              =================================================================================
setting                                  description
===========================              =================================================================================
/builtdock/controller/host                    IP address of the host running controller
/builtdock/controller/port                    port used by the controller service (default: 8000)
/builtdock/controller/protocol                protocol for controller (default: http)
/builtdock/controller/secretKey               used for secrets (default: randomly generated)
/builtdock/controller/builderKey              used by builder to authenticate with the controller (default: randomly generated)
/builtdock/builder/users/*                    stores user SSH keys (used by builder)
/builtdock/domains/*                          domain configuration for applications (used by router)
===========================              =================================================================================

Settings used by controller
---------------------------
The following etcd keys are used by the controller component.

====================================      ======================================================
setting                                   description
====================================      ======================================================
/builtdock/controller/registrationEnabled      enable registration for new Deis users (default: true)
/builtdock/controller/webEnabled               enable controller web UI (default: false)
/builtdock/cache/host                          host of the cache component (set by cache)
/builtdock/cache/port                          port of the cache component (set by cache)
/builtdock/database/host                       host of the database component (set by database)
/builtdock/database/port                       port of the database component (set by database)
/builtdock/database/engine                     database engine (set by database)
/builtdock/database/name                       database name (set by database)
/builtdock/database/user                       database user (set by database)
/builtdock/database/password                   database password (set by database)
/builtdock/registry/host                       host of the registry component (set by registry)
/builtdock/registry/port                       port of the registry component (set by registry)
/builtdock/registry/protocol                   protocol of the registry component (set by registry)
====================================      ======================================================

Using a custom controller image
-------------------------------
You can use a custom Docker image for the controller component instead of the image
supplied with Deis:

.. code-block:: console

    $ etcdctl set /builtdock/controller/image myaccount/myimage:latest

This will pull the image from the public Docker registry. You can also pull from a private
registry:

.. code-block:: console

    $ etcdctl set /builtdock/controller/image registry.mydomain.org:5000/myaccount/myimage:latest

Be sure that your custom image functions in the same way as the `stock controller image`_ shipped with
Deis. Specifically, ensure that it sets and reads appropriate etcd keys.

.. _`stock controller image`: https://github.com/builtdock/builtdock/tree/master/controller
.. _`#985`: https://github.com/builtdock/deis/issues/985

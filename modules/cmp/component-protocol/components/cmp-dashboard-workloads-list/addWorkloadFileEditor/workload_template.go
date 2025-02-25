// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package addWorkloadFileEditor

import "github.com/erda-project/erda/apistructs"

var workloadTemplates = map[string]string{
	string(apistructs.K8SDeployment): `apiVersion: apps/v1
kind: Deployment
metadata:
  name: #string
  annotations:
    {}
#    key: string
  labels:
    {}
#    key: string
spec:
  selector:
    matchLabels:
      erda/workloadselector: apps.deployment-default-#name
#      key: string
#    matchExpressions:
#      - key: string
#        operator: string
#        values:
#          - string
  template:
    metadata:
      labels:
        erda/workloadselector: apps.deployment-default-#name
#        key: string
#      annotations:
#        key: string
#      clusterName: string
#      creationTimestamp: string
#      deletionGracePeriodSeconds: int
#      deletionTimestamp: string
#      finalizers:
#        - string
#      generateName: string
#      generation: int
#      managedFields:
#        - apiVersion: string
#          fieldsType: string
#          fieldsV1: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
#          manager: string
#          operation: string
#          time: string
#      name: string
#      namespace: string
#      ownerReferences:
#        - apiVersion: string
#          blockOwnerDeletion: boolean
#          controller: boolean
#          kind: string
#          name: string
#          uid: string
#      resourceVersion: string
#      selfLink: string
#      uid: string
    spec:
      containers:
        - imagePullPolicy: Always
          name: container-0
          _active: true
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
      affinity:
#        nodeAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - preference:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchFields:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            nodeSelectorTerms:
#              - matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchFields:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#        podAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - podAffinityTerm:
#                labelSelector:
#                  matchExpressions:
#                    - key: string
#                      operator: string
#                      values:
#                        - string
#                  matchLabels:
#                    key: string
#                namespaces:
#                  - string
#                topologyKey: string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              namespaces:
#                - string
#              topologyKey: string
#        podAntiAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - podAffinityTerm:
#                labelSelector:
#                  matchExpressions:
#                    - key: string
#                      operator: string
#                      values:
#                        - string
#                  matchLabels:
#                    key: string
#                namespaces:
#                  - string
#                topologyKey: string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              namespaces:
#                - string
#              topologyKey: string
      imagePullSecrets:
#        - name: string
      initContainers:
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
      restartPolicy: Always
      volumes:
#        - awsElasticBlockStore:
#            fsType: string
#            partition: int
#            readOnly: boolean
#            volumeID: string
#          azureDisk:
#            cachingMode: string
#            diskName: string
#            diskURI: string
#            fsType: string
#            kind: string
#            readOnly: boolean
#          azureFile:
#            readOnly: boolean
#            secretName: string
#            shareName: string
#          cephfs:
#            monitors:
#              - string
#            path: string
#            readOnly: boolean
#            secretFile: string
#            secretRef:
#              name: string
#            user: string
#          cinder:
#            fsType: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            volumeID: string
#          configMap:
#            defaultMode: int
#            items:
#              - key: string
#                mode: int
#                path: string
#            name: string
#            optional: boolean
#          csi:
#            driver: string
#            fsType: string
#            nodePublishSecretRef:
#              name: string
#            readOnly: boolean
#            volumeAttributes:
#              key: string
#          downwardAPI:
#            defaultMode: int
#            items:
#              - fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                mode: int
#                path: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#          emptyDir:
#            medium: string
#            sizeLimit: string
#          fc:
#            fsType: string
#            lun: int
#            readOnly: boolean
#            targetWWNs:
#              - string
#            wwids:
#              - string
#          flexVolume:
#            driver: string
#            fsType: string
#            options:
#              key: string
#            readOnly: boolean
#            secretRef:
#              name: string
#          flocker:
#            datasetName: string
#            datasetUUID: string
#          gcePersistentDisk:
#            fsType: string
#            partition: int
#            pdName: string
#            readOnly: boolean
#          gitRepo:
#            directory: string
#            repository: string
#            revision: string
#          glusterfs:
#            endpoints: string
#            path: string
#            readOnly: boolean
#          hostPath:
#            type: string
#            path: string
#          iscsi:
#            chapAuthDiscovery: boolean
#            chapAuthSession: boolean
#            fsType: string
#            initiatorName: string
#            iqn: string
#            iscsiInterface: string
#            lun: int
#            portals:
#              - string
#            readOnly: boolean
#            secretRef:
#              name: string
#            targetPortal: string
#          name: string
#          nfs:
#            path: string
#            readOnly: boolean
#            server: string
#          persistentVolumeClaim:
#            claimName: string
#            readOnly: boolean
#          photonPersistentDisk:
#            fsType: string
#            pdID: string
#          portworxVolume:
#            fsType: string
#            readOnly: boolean
#            volumeID: string
#          projected:
#            defaultMode: int
#            sources:
#              - configMap:
#                  items:
#                    - key: string
#                      mode: int
#                      path: string
#                  name: string
#                  optional: boolean
#                downwardAPI:
#                  items:
#                    - fieldRef:
#                        apiVersion: string
#                        fieldPath: string
#                      mode: int
#                      path: string
#                      resourceFieldRef:
#                        containerName: string
#                        divisor: string
#                        resource: string
#                secret:
#                  items:
#                    - key: string
#                      mode: int
#                      path: string
#                  name: string
#                  optional: boolean
#                serviceAccountToken:
#                  audience: string
#                  expirationSeconds: int
#                  path: string
#          quobyte:
#            group: string
#            readOnly: boolean
#            registry: string
#            tenant: string
#            user: string
#            volume: string
#          rbd:
#            fsType: string
#            image: string
#            keyring: string
#            monitors:
#              - string
#            pool: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            user: string
#          scaleIO:
#            fsType: string
#            gateway: string
#            protectionDomain: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            sslEnabled: boolean
#            storageMode: string
#            storagePool: string
#            system: string
#            volumeName: string
#          secret:
#            defaultMode: int
#            items:
#              - key: string
#                mode: int
#                path: string
#            optional: boolean
#            secretName: string
#          storageos:
#            fsType: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            volumeName: string
#            volumeNamespace: string
#          vsphereVolume:
#            fsType: string
#            storagePolicyID: string
#            storagePolicyName: string
#            volumePath: string
#      activeDeadlineSeconds: int
#      automountServiceAccountToken: boolean
#      dnsConfig:
#        nameservers:
#          - string
#        options:
#          - name: string
#            value: string
#        searches:
#          - string
#      dnsPolicy: string
#      enableServiceLinks: boolean
#      ephemeralContainers:
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          targetContainerName: string
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
#      hostAliases:
#        - hostnames:
#            - string
#          ip: string
#      hostIPC: boolean
#      hostNetwork: boolean
#      hostPID: boolean
#      hostname: string
#      nodeName: string
#      nodeSelector:
#        key: string
#      overhead:
#        key: string
#      preemptionPolicy: string
#      priority: int
#      priorityClassName: string
#      readinessGates:
#        - conditionType: string
#      runtimeClassName: string
#      schedulerName: string
#      securityContext:
#        fsGroup: int
#        fsGroupChangePolicy: string
#        runAsGroup: int
#        runAsNonRoot: boolean
#        runAsUser: int
#        seLinuxOptions:
#          type: string
#          level: string
#          role: string
#          user: string
#        supplementalGroups:
#          - int
#        sysctls:
#          - name: string
#            value: string
#        windowsOptions:
#          gmsaCredentialSpec: string
#          gmsaCredentialSpecName: string
#          runAsUserName: string
#      serviceAccount: string
#      serviceAccountName: string
#      shareProcessNamespace: boolean
#      subdomain: string
#      terminationGracePeriodSeconds: int
#      tolerations:
#        - effect: string
#          key: string
#          operator: string
#          tolerationSeconds: int
#          value: string
#      topologySpreadConstraints:
#        - labelSelector:
#            matchExpressions:
#              - key: string
#                operator: string
#                values:
#                  - string
#            matchLabels:
#              key: string
#          maxSkew: int
#          topologyKey: string
#          whenUnsatisfiable: string
  replicas: 1
#  minReadySeconds: int
#  paused: boolean
#  progressDeadlineSeconds: int
#  revisionHistoryLimit: int
#  strategy:
#    type: string
#    rollingUpdate:
#      maxSurge: string
#      maxUnavailable: string
`,
	string(apistructs.K8SStatefulSet): `apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: #string
  annotations:
    {}
#    key: string
  labels:
    {}
#    key: string
spec:
  selector:
    matchLabels:
      erda/workloadselector: apps.statefulset-default-#name
#      key: string
#    matchExpressions:
#      - key: string
#        operator: string
#        values:
#          - string
  template:
    metadata:
      labels:
        erda/workloadselector: apps.statefulset-default-#name
#        key: string
#      annotations:
#        key: string
#      clusterName: string
#      creationTimestamp: string
#      deletionGracePeriodSeconds: int
#      deletionTimestamp: string
#      finalizers:
#        - string
#      generateName: string
#      generation: int
#      managedFields:
#        - apiVersion: string
#          fieldsType: string
#          fieldsV1: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
#          manager: string
#          operation: string
#          time: string
#      name: string
#      namespace: string
#      ownerReferences:
#        - apiVersion: string
#          blockOwnerDeletion: boolean
#          controller: boolean
#          kind: string
#          name: string
#          uid: string
#      resourceVersion: string
#      selfLink: string
#      uid: string
    spec:
      containers:
        - imagePullPolicy: Always
          name: container-0
          _active: true
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
      affinity:
#        nodeAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - preference:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchFields:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            nodeSelectorTerms:
#              - matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchFields:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#        podAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - podAffinityTerm:
#                labelSelector:
#                  matchExpressions:
#                    - key: string
#                      operator: string
#                      values:
#                        - string
#                  matchLabels:
#                    key: string
#                namespaces:
#                  - string
#                topologyKey: string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              namespaces:
#                - string
#              topologyKey: string
#        podAntiAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - podAffinityTerm:
#                labelSelector:
#                  matchExpressions:
#                    - key: string
#                      operator: string
#                      values:
#                        - string
#                  matchLabels:
#                    key: string
#                namespaces:
#                  - string
#                topologyKey: string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              namespaces:
#                - string
#              topologyKey: string
      imagePullSecrets:
#        - name: string
      initContainers:
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
      restartPolicy: Always
      volumes:
#        - awsElasticBlockStore:
#            fsType: string
#            partition: int
#            readOnly: boolean
#            volumeID: string
#          azureDisk:
#            cachingMode: string
#            diskName: string
#            diskURI: string
#            fsType: string
#            kind: string
#            readOnly: boolean
#          azureFile:
#            readOnly: boolean
#            secretName: string
#            shareName: string
#          cephfs:
#            monitors:
#              - string
#            path: string
#            readOnly: boolean
#            secretFile: string
#            secretRef:
#              name: string
#            user: string
#          cinder:
#            fsType: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            volumeID: string
#          configMap:
#            defaultMode: int
#            items:
#              - key: string
#                mode: int
#                path: string
#            name: string
#            optional: boolean
#          csi:
#            driver: string
#            fsType: string
#            nodePublishSecretRef:
#              name: string
#            readOnly: boolean
#            volumeAttributes:
#              key: string
#          downwardAPI:
#            defaultMode: int
#            items:
#              - fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                mode: int
#                path: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#          emptyDir:
#            medium: string
#            sizeLimit: string
#          fc:
#            fsType: string
#            lun: int
#            readOnly: boolean
#            targetWWNs:
#              - string
#            wwids:
#              - string
#          flexVolume:
#            driver: string
#            fsType: string
#            options:
#              key: string
#            readOnly: boolean
#            secretRef:
#              name: string
#          flocker:
#            datasetName: string
#            datasetUUID: string
#          gcePersistentDisk:
#            fsType: string
#            partition: int
#            pdName: string
#            readOnly: boolean
#          gitRepo:
#            directory: string
#            repository: string
#            revision: string
#          glusterfs:
#            endpoints: string
#            path: string
#            readOnly: boolean
#          hostPath:
#            type: string
#            path: string
#          iscsi:
#            chapAuthDiscovery: boolean
#            chapAuthSession: boolean
#            fsType: string
#            initiatorName: string
#            iqn: string
#            iscsiInterface: string
#            lun: int
#            portals:
#              - string
#            readOnly: boolean
#            secretRef:
#              name: string
#            targetPortal: string
#          name: string
#          nfs:
#            path: string
#            readOnly: boolean
#            server: string
#          persistentVolumeClaim:
#            claimName: string
#            readOnly: boolean
#          photonPersistentDisk:
#            fsType: string
#            pdID: string
#          portworxVolume:
#            fsType: string
#            readOnly: boolean
#            volumeID: string
#          projected:
#            defaultMode: int
#            sources:
#              - configMap:
#                  items:
#                    - key: string
#                      mode: int
#                      path: string
#                  name: string
#                  optional: boolean
#                downwardAPI:
#                  items:
#                    - fieldRef:
#                        apiVersion: string
#                        fieldPath: string
#                      mode: int
#                      path: string
#                      resourceFieldRef:
#                        containerName: string
#                        divisor: string
#                        resource: string
#                secret:
#                  items:
#                    - key: string
#                      mode: int
#                      path: string
#                  name: string
#                  optional: boolean
#                serviceAccountToken:
#                  audience: string
#                  expirationSeconds: int
#                  path: string
#          quobyte:
#            group: string
#            readOnly: boolean
#            registry: string
#            tenant: string
#            user: string
#            volume: string
#          rbd:
#            fsType: string
#            image: string
#            keyring: string
#            monitors:
#              - string
#            pool: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            user: string
#          scaleIO:
#            fsType: string
#            gateway: string
#            protectionDomain: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            sslEnabled: boolean
#            storageMode: string
#            storagePool: string
#            system: string
#            volumeName: string
#          secret:
#            defaultMode: int
#            items:
#              - key: string
#                mode: int
#                path: string
#            optional: boolean
#            secretName: string
#          storageos:
#            fsType: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            volumeName: string
#            volumeNamespace: string
#          vsphereVolume:
#            fsType: string
#            storagePolicyID: string
#            storagePolicyName: string
#            volumePath: string
#      activeDeadlineSeconds: int
#      automountServiceAccountToken: boolean
#      dnsConfig:
#        nameservers:
#          - string
#        options:
#          - name: string
#            value: string
#        searches:
#          - string
#      dnsPolicy: string
#      enableServiceLinks: boolean
#      ephemeralContainers:
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          targetContainerName: string
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
#      hostAliases:
#        - hostnames:
#            - string
#          ip: string
#      hostIPC: boolean
#      hostNetwork: boolean
#      hostPID: boolean
#      hostname: string
#      nodeName: string
#      nodeSelector:
#        key: string
#      overhead:
#        key: string
#      preemptionPolicy: string
#      priority: int
#      priorityClassName: string
#      readinessGates:
#        - conditionType: string
#      runtimeClassName: string
#      schedulerName: string
#      securityContext:
#        fsGroup: int
#        fsGroupChangePolicy: string
#        runAsGroup: int
#        runAsNonRoot: boolean
#        runAsUser: int
#        seLinuxOptions:
#          type: string
#          level: string
#          role: string
#          user: string
#        supplementalGroups:
#          - int
#        sysctls:
#          - name: string
#            value: string
#        windowsOptions:
#          gmsaCredentialSpec: string
#          gmsaCredentialSpecName: string
#          runAsUserName: string
#      serviceAccount: string
#      serviceAccountName: string
#      shareProcessNamespace: boolean
#      subdomain: string
#      terminationGracePeriodSeconds: int
#      tolerations:
#        - effect: string
#          key: string
#          operator: string
#          tolerationSeconds: int
#          value: string
#      topologySpreadConstraints:
#        - labelSelector:
#            matchExpressions:
#              - key: string
#                operator: string
#                values:
#                  - string
#            matchLabels:
#              key: string
#          maxSkew: int
#          topologyKey: string
#          whenUnsatisfiable: string
  replicas: 1
#  podManagementPolicy: string
#  revisionHistoryLimit: int
#  serviceName: string
#  updateStrategy:
#    type: string
#    rollingUpdate:
#      partition: int
#  volumeClaimTemplates:
#    - Error loading schema for io.k8s.api.core.v1.PersistentVolumeClaim
`,
	string(apistructs.K8SDaemonSet): `apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: #string
  annotations:
    {}
#    key: string
  labels:
    {}
#    key: string
spec:
  selector:
    matchLabels:
      erda/workloadselector: apps.daemonset-default-#name
#      key: string
#    matchExpressions:
#      - key: string
#        operator: string
#        values:
#          - string
  template:
    metadata:
      labels:
        erda/workloadselector: apps.daemonset-default-#name
#        key: string
#      annotations:
#        key: string
#      clusterName: string
#      creationTimestamp: string
#      deletionGracePeriodSeconds: int
#      deletionTimestamp: string
#      finalizers:
#        - string
#      generateName: string
#      generation: int
#      managedFields:
#        - apiVersion: string
#          fieldsType: string
#          fieldsV1: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
#          manager: string
#          operation: string
#          time: string
#      name: string
#      namespace: string
#      ownerReferences:
#        - apiVersion: string
#          blockOwnerDeletion: boolean
#          controller: boolean
#          kind: string
#          name: string
#          uid: string
#      resourceVersion: string
#      selfLink: string
#      uid: string
    spec:
      containers:
        - imagePullPolicy: Always
          name: container-0
          _active: true
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
      affinity:
#        nodeAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - preference:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchFields:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            nodeSelectorTerms:
#              - matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchFields:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#        podAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - podAffinityTerm:
#                labelSelector:
#                  matchExpressions:
#                    - key: string
#                      operator: string
#                      values:
#                        - string
#                  matchLabels:
#                    key: string
#                namespaces:
#                  - string
#                topologyKey: string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              namespaces:
#                - string
#              topologyKey: string
#        podAntiAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - podAffinityTerm:
#                labelSelector:
#                  matchExpressions:
#                    - key: string
#                      operator: string
#                      values:
#                        - string
#                  matchLabels:
#                    key: string
#                namespaces:
#                  - string
#                topologyKey: string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              namespaces:
#                - string
#              topologyKey: string
      imagePullSecrets:
#        - name: string
      initContainers:
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
      restartPolicy: Always
      volumes:
#        - awsElasticBlockStore:
#            fsType: string
#            partition: int
#            readOnly: boolean
#            volumeID: string
#          azureDisk:
#            cachingMode: string
#            diskName: string
#            diskURI: string
#            fsType: string
#            kind: string
#            readOnly: boolean
#          azureFile:
#            readOnly: boolean
#            secretName: string
#            shareName: string
#          cephfs:
#            monitors:
#              - string
#            path: string
#            readOnly: boolean
#            secretFile: string
#            secretRef:
#              name: string
#            user: string
#          cinder:
#            fsType: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            volumeID: string
#          configMap:
#            defaultMode: int
#            items:
#              - key: string
#                mode: int
#                path: string
#            name: string
#            optional: boolean
#          csi:
#            driver: string
#            fsType: string
#            nodePublishSecretRef:
#              name: string
#            readOnly: boolean
#            volumeAttributes:
#              key: string
#          downwardAPI:
#            defaultMode: int
#            items:
#              - fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                mode: int
#                path: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#          emptyDir:
#            medium: string
#            sizeLimit: string
#          fc:
#            fsType: string
#            lun: int
#            readOnly: boolean
#            targetWWNs:
#              - string
#            wwids:
#              - string
#          flexVolume:
#            driver: string
#            fsType: string
#            options:
#              key: string
#            readOnly: boolean
#            secretRef:
#              name: string
#          flocker:
#            datasetName: string
#            datasetUUID: string
#          gcePersistentDisk:
#            fsType: string
#            partition: int
#            pdName: string
#            readOnly: boolean
#          gitRepo:
#            directory: string
#            repository: string
#            revision: string
#          glusterfs:
#            endpoints: string
#            path: string
#            readOnly: boolean
#          hostPath:
#            type: string
#            path: string
#          iscsi:
#            chapAuthDiscovery: boolean
#            chapAuthSession: boolean
#            fsType: string
#            initiatorName: string
#            iqn: string
#            iscsiInterface: string
#            lun: int
#            portals:
#              - string
#            readOnly: boolean
#            secretRef:
#              name: string
#            targetPortal: string
#          name: string
#          nfs:
#            path: string
#            readOnly: boolean
#            server: string
#          persistentVolumeClaim:
#            claimName: string
#            readOnly: boolean
#          photonPersistentDisk:
#            fsType: string
#            pdID: string
#          portworxVolume:
#            fsType: string
#            readOnly: boolean
#            volumeID: string
#          projected:
#            defaultMode: int
#            sources:
#              - configMap:
#                  items:
#                    - key: string
#                      mode: int
#                      path: string
#                  name: string
#                  optional: boolean
#                downwardAPI:
#                  items:
#                    - fieldRef:
#                        apiVersion: string
#                        fieldPath: string
#                      mode: int
#                      path: string
#                      resourceFieldRef:
#                        containerName: string
#                        divisor: string
#                        resource: string
#                secret:
#                  items:
#                    - key: string
#                      mode: int
#                      path: string
#                  name: string
#                  optional: boolean
#                serviceAccountToken:
#                  audience: string
#                  expirationSeconds: int
#                  path: string
#          quobyte:
#            group: string
#            readOnly: boolean
#            registry: string
#            tenant: string
#            user: string
#            volume: string
#          rbd:
#            fsType: string
#            image: string
#            keyring: string
#            monitors:
#              - string
#            pool: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            user: string
#          scaleIO:
#            fsType: string
#            gateway: string
#            protectionDomain: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            sslEnabled: boolean
#            storageMode: string
#            storagePool: string
#            system: string
#            volumeName: string
#          secret:
#            defaultMode: int
#            items:
#              - key: string
#                mode: int
#                path: string
#            optional: boolean
#            secretName: string
#          storageos:
#            fsType: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            volumeName: string
#            volumeNamespace: string
#          vsphereVolume:
#            fsType: string
#            storagePolicyID: string
#            storagePolicyName: string
#            volumePath: string
#      activeDeadlineSeconds: int
#      automountServiceAccountToken: boolean
#      dnsConfig:
#        nameservers:
#          - string
#        options:
#          - name: string
#            value: string
#        searches:
#          - string
#      dnsPolicy: string
#      enableServiceLinks: boolean
#      ephemeralContainers:
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          targetContainerName: string
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
#      hostAliases:
#        - hostnames:
#            - string
#          ip: string
#      hostIPC: boolean
#      hostNetwork: boolean
#      hostPID: boolean
#      hostname: string
#      nodeName: string
#      nodeSelector:
#        key: string
#      overhead:
#        key: string
#      preemptionPolicy: string
#      priority: int
#      priorityClassName: string
#      readinessGates:
#        - conditionType: string
#      runtimeClassName: string
#      schedulerName: string
#      securityContext:
#        fsGroup: int
#        fsGroupChangePolicy: string
#        runAsGroup: int
#        runAsNonRoot: boolean
#        runAsUser: int
#        seLinuxOptions:
#          type: string
#          level: string
#          role: string
#          user: string
#        supplementalGroups:
#          - int
#        sysctls:
#          - name: string
#            value: string
#        windowsOptions:
#          gmsaCredentialSpec: string
#          gmsaCredentialSpecName: string
#          runAsUserName: string
#      serviceAccount: string
#      serviceAccountName: string
#      shareProcessNamespace: boolean
#      subdomain: string
#      terminationGracePeriodSeconds: int
#      tolerations:
#        - effect: string
#          key: string
#          operator: string
#          tolerationSeconds: int
#          value: string
#      topologySpreadConstraints:
#        - labelSelector:
#            matchExpressions:
#              - key: string
#                operator: string
#                values:
#                  - string
#            matchLabels:
#              key: string
#          maxSkew: int
#          topologyKey: string
#          whenUnsatisfiable: string
#  minReadySeconds: int
#  revisionHistoryLimit: int
#  updateStrategy:
#    type: string
#    rollingUpdate:
#      maxUnavailable: string
`,
	string(apistructs.K8SJob): `apiVersion: batch/v1
kind: Job
metadata:
  name: #string
  annotations:
    {}
#    key: string
  labels:
    {}
#    key: string
spec:
  selector:
    matchLabels:
#      key: string
#    matchExpressions:
#      - key: string
#        operator: string
#        values:
#          - string
  template:
    metadata:
      labels:
        {}
#        key: string
#      annotations:
#        key: string
#      clusterName: string
#      creationTimestamp: string
#      deletionGracePeriodSeconds: int
#      deletionTimestamp: string
#      finalizers:
#        - string
#      generateName: string
#      generation: int
#      managedFields:
#        - apiVersion: string
#          fieldsType: string
#          fieldsV1: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
#          manager: string
#          operation: string
#          time: string
#      name: string
#      namespace: string
#      ownerReferences:
#        - apiVersion: string
#          blockOwnerDeletion: boolean
#          controller: boolean
#          kind: string
#          name: string
#          uid: string
#      resourceVersion: string
#      selfLink: string
#      uid: string
    spec:
      containers:
        - imagePullPolicy: Always
          name: container-0
          _active: true
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
      affinity:
#        nodeAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - preference:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchFields:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            nodeSelectorTerms:
#              - matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchFields:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#        podAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - podAffinityTerm:
#                labelSelector:
#                  matchExpressions:
#                    - key: string
#                      operator: string
#                      values:
#                        - string
#                  matchLabels:
#                    key: string
#                namespaces:
#                  - string
#                topologyKey: string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              namespaces:
#                - string
#              topologyKey: string
#        podAntiAffinity:
#          preferredDuringSchedulingIgnoredDuringExecution:
#            - podAffinityTerm:
#                labelSelector:
#                  matchExpressions:
#                    - key: string
#                      operator: string
#                      values:
#                        - string
#                  matchLabels:
#                    key: string
#                namespaces:
#                  - string
#                topologyKey: string
#              weight: int
#          requiredDuringSchedulingIgnoredDuringExecution:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              namespaces:
#                - string
#              topologyKey: string
      imagePullSecrets:
#        - name: string
      initContainers:
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
      restartPolicy: Never
      volumes:
#        - awsElasticBlockStore:
#            fsType: string
#            partition: int
#            readOnly: boolean
#            volumeID: string
#          azureDisk:
#            cachingMode: string
#            diskName: string
#            diskURI: string
#            fsType: string
#            kind: string
#            readOnly: boolean
#          azureFile:
#            readOnly: boolean
#            secretName: string
#            shareName: string
#          cephfs:
#            monitors:
#              - string
#            path: string
#            readOnly: boolean
#            secretFile: string
#            secretRef:
#              name: string
#            user: string
#          cinder:
#            fsType: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            volumeID: string
#          configMap:
#            defaultMode: int
#            items:
#              - key: string
#                mode: int
#                path: string
#            name: string
#            optional: boolean
#          csi:
#            driver: string
#            fsType: string
#            nodePublishSecretRef:
#              name: string
#            readOnly: boolean
#            volumeAttributes:
#              key: string
#          downwardAPI:
#            defaultMode: int
#            items:
#              - fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                mode: int
#                path: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#          emptyDir:
#            medium: string
#            sizeLimit: string
#          fc:
#            fsType: string
#            lun: int
#            readOnly: boolean
#            targetWWNs:
#              - string
#            wwids:
#              - string
#          flexVolume:
#            driver: string
#            fsType: string
#            options:
#              key: string
#            readOnly: boolean
#            secretRef:
#              name: string
#          flocker:
#            datasetName: string
#            datasetUUID: string
#          gcePersistentDisk:
#            fsType: string
#            partition: int
#            pdName: string
#            readOnly: boolean
#          gitRepo:
#            directory: string
#            repository: string
#            revision: string
#          glusterfs:
#            endpoints: string
#            path: string
#            readOnly: boolean
#          hostPath:
#            type: string
#            path: string
#          iscsi:
#            chapAuthDiscovery: boolean
#            chapAuthSession: boolean
#            fsType: string
#            initiatorName: string
#            iqn: string
#            iscsiInterface: string
#            lun: int
#            portals:
#              - string
#            readOnly: boolean
#            secretRef:
#              name: string
#            targetPortal: string
#          name: string
#          nfs:
#            path: string
#            readOnly: boolean
#            server: string
#          persistentVolumeClaim:
#            claimName: string
#            readOnly: boolean
#          photonPersistentDisk:
#            fsType: string
#            pdID: string
#          portworxVolume:
#            fsType: string
#            readOnly: boolean
#            volumeID: string
#          projected:
#            defaultMode: int
#            sources:
#              - configMap:
#                  items:
#                    - key: string
#                      mode: int
#                      path: string
#                  name: string
#                  optional: boolean
#                downwardAPI:
#                  items:
#                    - fieldRef:
#                        apiVersion: string
#                        fieldPath: string
#                      mode: int
#                      path: string
#                      resourceFieldRef:
#                        containerName: string
#                        divisor: string
#                        resource: string
#                secret:
#                  items:
#                    - key: string
#                      mode: int
#                      path: string
#                  name: string
#                  optional: boolean
#                serviceAccountToken:
#                  audience: string
#                  expirationSeconds: int
#                  path: string
#          quobyte:
#            group: string
#            readOnly: boolean
#            registry: string
#            tenant: string
#            user: string
#            volume: string
#          rbd:
#            fsType: string
#            image: string
#            keyring: string
#            monitors:
#              - string
#            pool: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            user: string
#          scaleIO:
#            fsType: string
#            gateway: string
#            protectionDomain: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            sslEnabled: boolean
#            storageMode: string
#            storagePool: string
#            system: string
#            volumeName: string
#          secret:
#            defaultMode: int
#            items:
#              - key: string
#                mode: int
#                path: string
#            optional: boolean
#            secretName: string
#          storageos:
#            fsType: string
#            readOnly: boolean
#            secretRef:
#              name: string
#            volumeName: string
#            volumeNamespace: string
#          vsphereVolume:
#            fsType: string
#            storagePolicyID: string
#            storagePolicyName: string
#            volumePath: string
#      activeDeadlineSeconds: int
#      automountServiceAccountToken: boolean
#      dnsConfig:
#        nameservers:
#          - string
#        options:
#          - name: string
#            value: string
#        searches:
#          - string
#      dnsPolicy: string
#      enableServiceLinks: boolean
#      ephemeralContainers:
#        - args:
#            - string
#          command:
#            - string
#          env:
#            - name: string
#              value: string
#              valueFrom:
#                configMapKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#                fieldRef:
#                  apiVersion: string
#                  fieldPath: string
#                resourceFieldRef:
#                  containerName: string
#                  divisor: string
#                  resource: string
#                secretKeyRef:
#                  key: string
#                  name: string
#                  optional: boolean
#          envFrom:
#            - configMapRef:
#                name: string
#                optional: boolean
#              prefix: string
#              secretRef:
#                name: string
#                optional: boolean
#          image: string
#          imagePullPolicy: string
#          lifecycle:
#            postStart:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#            preStop:
#              exec:
#                command:
#                  - string
#              httpGet:
#                host: string
#                httpHeaders:
#                  - name: string
#                    value: string
#                path: string
#                port: string
#                scheme: string
#              tcpSocket:
#                host: string
#                port: string
#          livenessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          name: string
#          ports:
#            - containerPort: int
#              hostIP: string
#              hostPort: int
#              name: string
#              protocol: string
#          readinessProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          resources:
#            limits:
#              key: string
#            requests:
#              key: string
#          securityContext:
#            allowPrivilegeEscalation: boolean
#            capabilities:
#              add:
#                - string
#              drop:
#                - string
#            privileged: boolean
#            procMount: string
#            readOnlyRootFilesystem: boolean
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          startupProbe:
#            exec:
#              command:
#                - string
#            failureThreshold: int
#            httpGet:
#              host: string
#              httpHeaders:
#                - name: string
#                  value: string
#              path: string
#              port: string
#              scheme: string
#            initialDelaySeconds: int
#            periodSeconds: int
#            successThreshold: int
#            tcpSocket:
#              host: string
#              port: string
#            timeoutSeconds: int
#          stdin: boolean
#          stdinOnce: boolean
#          targetContainerName: string
#          terminationMessagePath: string
#          terminationMessagePolicy: string
#          tty: boolean
#          volumeDevices:
#            - devicePath: string
#              name: string
#          volumeMounts:
#            - mountPath: string
#              mountPropagation: string
#              name: string
#              readOnly: boolean
#              subPath: string
#              subPathExpr: string
#          workingDir: string
#      hostAliases:
#        - hostnames:
#            - string
#          ip: string
#      hostIPC: boolean
#      hostNetwork: boolean
#      hostPID: boolean
#      hostname: string
#      nodeName: string
#      nodeSelector:
#        key: string
#      overhead:
#        key: string
#      preemptionPolicy: string
#      priority: int
#      priorityClassName: string
#      readinessGates:
#        - conditionType: string
#      runtimeClassName: string
#      schedulerName: string
#      securityContext:
#        fsGroup: int
#        fsGroupChangePolicy: string
#        runAsGroup: int
#        runAsNonRoot: boolean
#        runAsUser: int
#        seLinuxOptions:
#          type: string
#          level: string
#          role: string
#          user: string
#        supplementalGroups:
#          - int
#        sysctls:
#          - name: string
#            value: string
#        windowsOptions:
#          gmsaCredentialSpec: string
#          gmsaCredentialSpecName: string
#          runAsUserName: string
#      serviceAccount: string
#      serviceAccountName: string
#      shareProcessNamespace: boolean
#      subdomain: string
#      terminationGracePeriodSeconds: int
#      tolerations:
#        - effect: string
#          key: string
#          operator: string
#          tolerationSeconds: int
#          value: string
#      topologySpreadConstraints:
#        - labelSelector:
#            matchExpressions:
#              - key: string
#                operator: string
#                values:
#                  - string
#            matchLabels:
#              key: string
#          maxSkew: int
#          topologyKey: string
#          whenUnsatisfiable: string
#  activeDeadlineSeconds: int
#  backoffLimit: int
#  completions: int
#  manualSelector: boolean
#  parallelism: int
#  ttlSecondsAfterFinished: int
`,
	string(apistructs.K8SCronJob): `apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: #string
  annotations:
    {}
#    key: string
  labels:
    {}
#    key: string
spec:
  jobTemplate:
    metadata:
      labels:
        {}
#        key: string
#      annotations:
#        key: string
#      clusterName: string
#      creationTimestamp: string
#      deletionGracePeriodSeconds: int
#      deletionTimestamp: string
#      finalizers:
#        - string
#      generateName: string
#      generation: int
#      managedFields:
#        - apiVersion: string
#          fieldsType: string
#          fieldsV1: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
#          manager: string
#          operation: string
#          time: string
#      name: string
#      namespace: string
#      ownerReferences:
#        - apiVersion: string
#          blockOwnerDeletion: boolean
#          controller: boolean
#          kind: string
#          name: string
#          uid: string
#      resourceVersion: string
#      selfLink: string
#      uid: string
    spec:
      template:
        spec:
          affinity:
#            nodeAffinity:
#              preferredDuringSchedulingIgnoredDuringExecution:
#                - preference:
#                    matchExpressions:
#                      - key: string
#                        operator: string
#                        values:
#                          - string
#                    matchFields:
#                      - key: string
#                        operator: string
#                        values:
#                          - string
#                  weight: int
#              requiredDuringSchedulingIgnoredDuringExecution:
#                nodeSelectorTerms:
#                  - matchExpressions:
#                      - key: string
#                        operator: string
#                        values:
#                          - string
#                    matchFields:
#                      - key: string
#                        operator: string
#                        values:
#                          - string
#            podAffinity:
#              preferredDuringSchedulingIgnoredDuringExecution:
#                - podAffinityTerm:
#                    labelSelector:
#                      matchExpressions:
#                        - key: string
#                          operator: string
#                          values:
#                            - string
#                      matchLabels:
#                        key: string
#                    namespaces:
#                      - string
#                    topologyKey: string
#                  weight: int
#              requiredDuringSchedulingIgnoredDuringExecution:
#                - labelSelector:
#                    matchExpressions:
#                      - key: string
#                        operator: string
#                        values:
#                          - string
#                    matchLabels:
#                      key: string
#                  namespaces:
#                    - string
#                  topologyKey: string
#            podAntiAffinity:
#              preferredDuringSchedulingIgnoredDuringExecution:
#                - podAffinityTerm:
#                    labelSelector:
#                      matchExpressions:
#                        - key: string
#                          operator: string
#                          values:
#                            - string
#                      matchLabels:
#                        key: string
#                    namespaces:
#                      - string
#                    topologyKey: string
#                  weight: int
#              requiredDuringSchedulingIgnoredDuringExecution:
#                - labelSelector:
#                    matchExpressions:
#                      - key: string
#                        operator: string
#                        values:
#                          - string
#                    matchLabels:
#                      key: string
#                  namespaces:
#                    - string
#                  topologyKey: string
          containers:
            - imagePullPolicy: Always
              name: container-0
              _active: true
#            - args:
#                - string
#              command:
#                - string
#              env:
#                - name: string
#                  value: string
#                  valueFrom:
#                    configMapKeyRef:
#                      key: string
#                      name: string
#                      optional: boolean
#                    fieldRef:
#                      apiVersion: string
#                      fieldPath: string
#                    resourceFieldRef:
#                      containerName: string
#                      divisor: string
#                      resource: string
#                    secretKeyRef:
#                      key: string
#                      name: string
#                      optional: boolean
#              envFrom:
#                - configMapRef:
#                    name: string
#                    optional: boolean
#                  prefix: string
#                  secretRef:
#                    name: string
#                    optional: boolean
#              image: string
#              imagePullPolicy: string
#              lifecycle:
#                postStart:
#                  exec:
#                    command:
#                      - string
#                  httpGet:
#                    host: string
#                    httpHeaders:
#                      - name: string
#                        value: string
#                    path: string
#                    port: string
#                    scheme: string
#                  tcpSocket:
#                    host: string
#                    port: string
#                preStop:
#                  exec:
#                    command:
#                      - string
#                  httpGet:
#                    host: string
#                    httpHeaders:
#                      - name: string
#                        value: string
#                    path: string
#                    port: string
#                    scheme: string
#                  tcpSocket:
#                    host: string
#                    port: string
#              livenessProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              name: string
#              ports:
#                - containerPort: int
#                  hostIP: string
#                  hostPort: int
#                  name: string
#                  protocol: string
#              readinessProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              resources:
#                limits:
#                  key: string
#                requests:
#                  key: string
#              securityContext:
#                allowPrivilegeEscalation: boolean
#                capabilities:
#                  add:
#                    - string
#                  drop:
#                    - string
#                privileged: boolean
#                procMount: string
#                readOnlyRootFilesystem: boolean
#                runAsGroup: int
#                runAsNonRoot: boolean
#                runAsUser: int
#                seLinuxOptions:
#                  type: string
#                  level: string
#                  role: string
#                  user: string
#                windowsOptions:
#                  gmsaCredentialSpec: string
#                  gmsaCredentialSpecName: string
#                  runAsUserName: string
#              startupProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              stdin: boolean
#              stdinOnce: boolean
#              terminationMessagePath: string
#              terminationMessagePolicy: string
#              tty: boolean
#              volumeDevices:
#                - devicePath: string
#                  name: string
#              volumeMounts:
#                - mountPath: string
#                  mountPropagation: string
#                  name: string
#                  readOnly: boolean
#                  subPath: string
#                  subPathExpr: string
#              workingDir: string
          imagePullSecrets:
#            - name: string
          initContainers:
#            - args:
#                - string
#              command:
#                - string
#              env:
#                - name: string
#                  value: string
#                  valueFrom:
#                    configMapKeyRef:
#                      key: string
#                      name: string
#                      optional: boolean
#                    fieldRef:
#                      apiVersion: string
#                      fieldPath: string
#                    resourceFieldRef:
#                      containerName: string
#                      divisor: string
#                      resource: string
#                    secretKeyRef:
#                      key: string
#                      name: string
#                      optional: boolean
#              envFrom:
#                - configMapRef:
#                    name: string
#                    optional: boolean
#                  prefix: string
#                  secretRef:
#                    name: string
#                    optional: boolean
#              image: string
#              imagePullPolicy: string
#              lifecycle:
#                postStart:
#                  exec:
#                    command:
#                      - string
#                  httpGet:
#                    host: string
#                    httpHeaders:
#                      - name: string
#                        value: string
#                    path: string
#                    port: string
#                    scheme: string
#                  tcpSocket:
#                    host: string
#                    port: string
#                preStop:
#                  exec:
#                    command:
#                      - string
#                  httpGet:
#                    host: string
#                    httpHeaders:
#                      - name: string
#                        value: string
#                    path: string
#                    port: string
#                    scheme: string
#                  tcpSocket:
#                    host: string
#                    port: string
#              livenessProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              name: string
#              ports:
#                - containerPort: int
#                  hostIP: string
#                  hostPort: int
#                  name: string
#                  protocol: string
#              readinessProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              resources:
#                limits:
#                  key: string
#                requests:
#                  key: string
#              securityContext:
#                allowPrivilegeEscalation: boolean
#                capabilities:
#                  add:
#                    - string
#                  drop:
#                    - string
#                privileged: boolean
#                procMount: string
#                readOnlyRootFilesystem: boolean
#                runAsGroup: int
#                runAsNonRoot: boolean
#                runAsUser: int
#                seLinuxOptions:
#                  type: string
#                  level: string
#                  role: string
#                  user: string
#                windowsOptions:
#                  gmsaCredentialSpec: string
#                  gmsaCredentialSpecName: string
#                  runAsUserName: string
#              startupProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              stdin: boolean
#              stdinOnce: boolean
#              terminationMessagePath: string
#              terminationMessagePolicy: string
#              tty: boolean
#              volumeDevices:
#                - devicePath: string
#                  name: string
#              volumeMounts:
#                - mountPath: string
#                  mountPropagation: string
#                  name: string
#                  readOnly: boolean
#                  subPath: string
#                  subPathExpr: string
#              workingDir: string
          restartPolicy: Never
          volumes:
#            - awsElasticBlockStore:
#                fsType: string
#                partition: int
#                readOnly: boolean
#                volumeID: string
#              azureDisk:
#                cachingMode: string
#                diskName: string
#                diskURI: string
#                fsType: string
#                kind: string
#                readOnly: boolean
#              azureFile:
#                readOnly: boolean
#                secretName: string
#                shareName: string
#              cephfs:
#                monitors:
#                  - string
#                path: string
#                readOnly: boolean
#                secretFile: string
#                secretRef:
#                  name: string
#                user: string
#              cinder:
#                fsType: string
#                readOnly: boolean
#                secretRef:
#                  name: string
#                volumeID: string
#              configMap:
#                defaultMode: int
#                items:
#                  - key: string
#                    mode: int
#                    path: string
#                name: string
#                optional: boolean
#              csi:
#                driver: string
#                fsType: string
#                nodePublishSecretRef:
#                  name: string
#                readOnly: boolean
#                volumeAttributes:
#                  key: string
#              downwardAPI:
#                defaultMode: int
#                items:
#                  - fieldRef:
#                      apiVersion: string
#                      fieldPath: string
#                    mode: int
#                    path: string
#                    resourceFieldRef:
#                      containerName: string
#                      divisor: string
#                      resource: string
#              emptyDir:
#                medium: string
#                sizeLimit: string
#              fc:
#                fsType: string
#                lun: int
#                readOnly: boolean
#                targetWWNs:
#                  - string
#                wwids:
#                  - string
#              flexVolume:
#                driver: string
#                fsType: string
#                options:
#                  key: string
#                readOnly: boolean
#                secretRef:
#                  name: string
#              flocker:
#                datasetName: string
#                datasetUUID: string
#              gcePersistentDisk:
#                fsType: string
#                partition: int
#                pdName: string
#                readOnly: boolean
#              gitRepo:
#                directory: string
#                repository: string
#                revision: string
#              glusterfs:
#                endpoints: string
#                path: string
#                readOnly: boolean
#              hostPath:
#                type: string
#                path: string
#              iscsi:
#                chapAuthDiscovery: boolean
#                chapAuthSession: boolean
#                fsType: string
#                initiatorName: string
#                iqn: string
#                iscsiInterface: string
#                lun: int
#                portals:
#                  - string
#                readOnly: boolean
#                secretRef:
#                  name: string
#                targetPortal: string
#              name: string
#              nfs:
#                path: string
#                readOnly: boolean
#                server: string
#              persistentVolumeClaim:
#                claimName: string
#                readOnly: boolean
#              photonPersistentDisk:
#                fsType: string
#                pdID: string
#              portworxVolume:
#                fsType: string
#                readOnly: boolean
#                volumeID: string
#              projected:
#                defaultMode: int
#                sources:
#                  - configMap:
#                      items:
#                        - key: string
#                          mode: int
#                          path: string
#                      name: string
#                      optional: boolean
#                    downwardAPI:
#                      items:
#                        - fieldRef:
#                            apiVersion: string
#                            fieldPath: string
#                          mode: int
#                          path: string
#                          resourceFieldRef:
#                            containerName: string
#                            divisor: string
#                            resource: string
#                    secret:
#                      items:
#                        - key: string
#                          mode: int
#                          path: string
#                      name: string
#                      optional: boolean
#                    serviceAccountToken:
#                      audience: string
#                      expirationSeconds: int
#                      path: string
#              quobyte:
#                group: string
#                readOnly: boolean
#                registry: string
#                tenant: string
#                user: string
#                volume: string
#              rbd:
#                fsType: string
#                image: string
#                keyring: string
#                monitors:
#                  - string
#                pool: string
#                readOnly: boolean
#                secretRef:
#                  name: string
#                user: string
#              scaleIO:
#                fsType: string
#                gateway: string
#                protectionDomain: string
#                readOnly: boolean
#                secretRef:
#                  name: string
#                sslEnabled: boolean
#                storageMode: string
#                storagePool: string
#                system: string
#                volumeName: string
#              secret:
#                defaultMode: int
#                items:
#                  - key: string
#                    mode: int
#                    path: string
#                optional: boolean
#                secretName: string
#              storageos:
#                fsType: string
#                readOnly: boolean
#                secretRef:
#                  name: string
#                volumeName: string
#                volumeNamespace: string
#              vsphereVolume:
#                fsType: string
#                storagePolicyID: string
#                storagePolicyName: string
#                volumePath: string
#          activeDeadlineSeconds: int
#          automountServiceAccountToken: boolean
#          dnsConfig:
#            nameservers:
#              - string
#            options:
#              - name: string
#                value: string
#            searches:
#              - string
#          dnsPolicy: string
#          enableServiceLinks: boolean
#          ephemeralContainers:
#            - args:
#                - string
#              command:
#                - string
#              env:
#                - name: string
#                  value: string
#                  valueFrom:
#                    configMapKeyRef:
#                      key: string
#                      name: string
#                      optional: boolean
#                    fieldRef:
#                      apiVersion: string
#                      fieldPath: string
#                    resourceFieldRef:
#                      containerName: string
#                      divisor: string
#                      resource: string
#                    secretKeyRef:
#                      key: string
#                      name: string
#                      optional: boolean
#              envFrom:
#                - configMapRef:
#                    name: string
#                    optional: boolean
#                  prefix: string
#                  secretRef:
#                    name: string
#                    optional: boolean
#              image: string
#              imagePullPolicy: string
#              lifecycle:
#                postStart:
#                  exec:
#                    command:
#                      - string
#                  httpGet:
#                    host: string
#                    httpHeaders:
#                      - name: string
#                        value: string
#                    path: string
#                    port: string
#                    scheme: string
#                  tcpSocket:
#                    host: string
#                    port: string
#                preStop:
#                  exec:
#                    command:
#                      - string
#                  httpGet:
#                    host: string
#                    httpHeaders:
#                      - name: string
#                        value: string
#                    path: string
#                    port: string
#                    scheme: string
#                  tcpSocket:
#                    host: string
#                    port: string
#              livenessProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              name: string
#              ports:
#                - containerPort: int
#                  hostIP: string
#                  hostPort: int
#                  name: string
#                  protocol: string
#              readinessProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              resources:
#                limits:
#                  key: string
#                requests:
#                  key: string
#              securityContext:
#                allowPrivilegeEscalation: boolean
#                capabilities:
#                  add:
#                    - string
#                  drop:
#                    - string
#                privileged: boolean
#                procMount: string
#                readOnlyRootFilesystem: boolean
#                runAsGroup: int
#                runAsNonRoot: boolean
#                runAsUser: int
#                seLinuxOptions:
#                  type: string
#                  level: string
#                  role: string
#                  user: string
#                windowsOptions:
#                  gmsaCredentialSpec: string
#                  gmsaCredentialSpecName: string
#                  runAsUserName: string
#              startupProbe:
#                exec:
#                  command:
#                    - string
#                failureThreshold: int
#                httpGet:
#                  host: string
#                  httpHeaders:
#                    - name: string
#                      value: string
#                  path: string
#                  port: string
#                  scheme: string
#                initialDelaySeconds: int
#                periodSeconds: int
#                successThreshold: int
#                tcpSocket:
#                  host: string
#                  port: string
#                timeoutSeconds: int
#              stdin: boolean
#              stdinOnce: boolean
#              targetContainerName: string
#              terminationMessagePath: string
#              terminationMessagePolicy: string
#              tty: boolean
#              volumeDevices:
#                - devicePath: string
#                  name: string
#              volumeMounts:
#                - mountPath: string
#                  mountPropagation: string
#                  name: string
#                  readOnly: boolean
#                  subPath: string
#                  subPathExpr: string
#              workingDir: string
#          hostAliases:
#            - hostnames:
#                - string
#              ip: string
#          hostIPC: boolean
#          hostNetwork: boolean
#          hostPID: boolean
#          hostname: string
#          nodeName: string
#          nodeSelector:
#            key: string
#          overhead:
#            key: string
#          preemptionPolicy: string
#          priority: int
#          priorityClassName: string
#          readinessGates:
#            - conditionType: string
#          runtimeClassName: string
#          schedulerName: string
#          securityContext:
#            fsGroup: int
#            fsGroupChangePolicy: string
#            runAsGroup: int
#            runAsNonRoot: boolean
#            runAsUser: int
#            seLinuxOptions:
#              type: string
#              level: string
#              role: string
#              user: string
#            supplementalGroups:
#              - int
#            sysctls:
#              - name: string
#                value: string
#            windowsOptions:
#              gmsaCredentialSpec: string
#              gmsaCredentialSpecName: string
#              runAsUserName: string
#          serviceAccount: string
#          serviceAccountName: string
#          shareProcessNamespace: boolean
#          subdomain: string
#          terminationGracePeriodSeconds: int
#          tolerations:
#            - effect: string
#              key: string
#              operator: string
#              tolerationSeconds: int
#              value: string
#          topologySpreadConstraints:
#            - labelSelector:
#                matchExpressions:
#                  - key: string
#                    operator: string
#                    values:
#                      - string
#                matchLabels:
#                  key: string
#              maxSkew: int
#              topologyKey: string
#              whenUnsatisfiable: string
#        metadata:
#          annotations:
#            key: string
#          clusterName: string
#          creationTimestamp: string
#          deletionGracePeriodSeconds: int
#          deletionTimestamp: string
#          finalizers:
#            - string
#          generateName: string
#          generation: int
#          labels:
#            key: string
#          managedFields:
#            - apiVersion: string
#              fieldsType: string
#              fieldsV1: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
#              manager: string
#              operation: string
#              time: string
#          name: string
#          namespace: string
#          ownerReferences:
#            - apiVersion: string
#              blockOwnerDeletion: boolean
#              controller: boolean
#              kind: string
#              name: string
#              uid: string
#          resourceVersion: string
#          selfLink: string
#          uid: string
#      activeDeadlineSeconds: int
#      backoffLimit: int
#      completions: int
#      manualSelector: boolean
#      parallelism: int
#      selector:
#        matchExpressions:
#          - key: string
#            operator: string
#            values:
#              - string
#        matchLabels:
#          key: string
#      ttlSecondsAfterFinished: int
#  concurrencyPolicy: string
#  failedJobsHistoryLimit: int
#  schedule: string
#  startingDeadlineSeconds: int
#  successfulJobsHistoryLimit: int
#  suspend: boolean
`,
}

import clsx from 'clsx';
import { BarChart, FileText, Info, Paperclip, Terminal } from 'lucide-react';

import { ContainerStatus } from '@/react/docker/containers/types';
import { Authorized } from '@/react/hooks/useUser';

import { Icon } from '@@/Icon';
import { Link } from '@@/Link';

import styles from './ContainerQuickActions.module.css';

export interface QuickActionsState {
  showQuickActionAttach: boolean;
  showQuickActionExec: boolean;
  showQuickActionInspect: boolean;
  showQuickActionLogs: boolean;
  showQuickActionStats: boolean;
}

export function ContainerQuickActions({
  status,
  containerId,
  nodeName,
  state,
}: {
  containerId: string;
  nodeName: string;
  status: ContainerStatus;
  state: QuickActionsState;
}) {
  const isActive =
    !!status &&
    [
      ContainerStatus.Starting,
      ContainerStatus.Running,
      ContainerStatus.Healthy,
      ContainerStatus.Unhealthy,
    ].includes(status);

  return (
    <div className={clsx('space-x-1', styles.root)}>
      {state.showQuickActionLogs && (
        <Authorized authorizations="DockerContainerLogs">
          <Link
            to="docker.containers.container.logs"
            params={{ id: containerId, nodeName }}
            title="Logs"
          >
            <Icon icon={FileText} className="space-right" />
          </Link>
        </Authorized>
      )}

      {state.showQuickActionInspect && (
        <Authorized authorizations="DockerContainerInspect">
          <Link
            to="docker.containers.container.inspect"
            params={{ id: containerId, nodeName }}
            title="Inspect"
          >
            <Icon icon={Info} className="space-right" />
          </Link>
        </Authorized>
      )}

      {state.showQuickActionStats && isActive && (
        <Authorized authorizations="DockerContainerStats">
          <Link
            to="docker.containers.container.stats"
            params={{ id: containerId, nodeName }}
            title="Stats"
          >
            <Icon icon={BarChart} className="space-right" />
          </Link>
        </Authorized>
      )}

      {state.showQuickActionExec && isActive && (
        <Authorized authorizations="DockerExecStart">
          <Link
            to="docker.containers.container.exec"
            params={{ id: containerId, nodeName }}
            title="Exec Console"
          >
            <Icon icon={Terminal} className="space-right" />
          </Link>
        </Authorized>
      )}

      {state.showQuickActionAttach && isActive && (
        <Authorized authorizations="DockerContainerAttach">
          <Link
            to="docker.containers.container.attach"
            params={{ id: containerId, nodeName }}
            title="Attach Console"
          >
            <Icon icon={Paperclip} className="space-right" />
          </Link>
        </Authorized>
      )}
    </div>
  );
}

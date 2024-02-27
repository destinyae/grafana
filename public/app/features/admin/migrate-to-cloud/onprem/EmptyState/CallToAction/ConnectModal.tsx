import React, { useId, useState } from 'react';

import { Modal, Button, Stack, TextLink, Field, Input, Text } from '@grafana/ui';
import { Trans, t } from 'app/core/internationalization';

interface Props {
  hideModal: () => void;
  onConfirm: (stackURL: string, token: string) => Promise<{ data: void } | { error: unknown }>;
}

export const ConnectModal = ({ hideModal, onConfirm }: Props) => {
  const [isConnecting, setIsConnecting] = useState(false);
  const [stackURL, setStackURL] = useState('');
  const [token, setToken] = useState('');
  const cloudStackId = useId();
  const tokenId = useId();

  const onConfirmConnect = async () => {
    setIsConnecting(true);
    await onConfirm(stackURL, token);
    setIsConnecting(false);
    hideModal();
  };

  return (
    <Modal isOpen title={t('migrate-to-cloud.connect-modal.title', 'Connect to a cloud stack')} onDismiss={hideModal}>
      <Text color="secondary">
        <Stack direction="column" gap={2}>
          <Trans i18nKey="migrate-to-cloud.connect-modal.body-get-started">
            To get started, you&apos;ll need a Grafana.com account.
          </Trans>
          <TextLink href="https://grafana.com/auth/sign-up/create-user?pg=prod-cloud" external>
            {t('migrate-to-cloud.connect-modal.body-sign-up', 'Sign up for a Grafana.com account')}
          </TextLink>
          <Trans i18nKey="migrate-to-cloud.connect-modal.body-cloud-stack">
            You&apos;ll also need a cloud stack. If you just signed up, we&apos;ll automatically create your first
            stack. If you have an account, you&apos;ll need to select or create a stack.
          </Trans>
          <TextLink href="https://grafana.com/auth/sign-in/" external>
            {t('migrate-to-cloud.connect-modal.body-view-stacks', 'View my cloud stacks')}
          </TextLink>
          <Trans i18nKey="migrate-to-cloud.connect-modal.body-paste-stack">
            Once you&apos;ve decided on a stack, paste the URL below.
          </Trans>
          <Field label={t('migrate-to-cloud.connect-modal.body-url-field', 'Cloud stack URL')}>
            <Input
              id={cloudStackId}
              value={stackURL}
              onChange={(event) => setStackURL(event.currentTarget.value)}
              placeholder="https://example.grafana.net/"
            />
          </Field>
          <span>
            <Trans i18nKey="migrate-to-cloud.connect-modal.body-token">
              Your self-managed Grafana installation needs special access to securely migrate content. You&apos;ll need
              to create a migration token on your chosen cloud stack.
            </Trans>
          </span>
          <span>
            <Trans i18nKey="migrate-to-cloud.connect-modal.body-token-instructions">
              Log into your cloud stack and navigate to Administration, General, Migrate to Grafana Cloud. Create a
              migration token on that screen and paste the token here.
            </Trans>
          </span>
          <Field label={t('migrate-to-cloud.connect-modal.body-token-field', 'Migration token')}>
            <Input
              id={tokenId}
              value={token}
              onChange={(event) => setToken(event.currentTarget.value)}
              placeholder={t('migrate-to-cloud.connect-modal.body-token-field-placeholder', 'Paste token here')}
            />
          </Field>
        </Stack>
      </Text>
      <Modal.ButtonRow>
        <Button variant="secondary" onClick={hideModal}>
          <Trans i18nKey="migrate-to-cloud.connect-modal.cancel">Cancel</Trans>
        </Button>
        <Button disabled={isConnecting || !(stackURL && token)} onClick={onConfirmConnect}>
          {isConnecting
            ? t('migrate-to-cloud.connect-modal.connecting', 'Connecting to this stack...')
            : t('migrate-to-cloud.connect-modal.connect', 'Connect to this stack')}
        </Button>
      </Modal.ButtonRow>
    </Modal>
  );
};

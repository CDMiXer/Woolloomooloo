// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* #5 shazhko04: перестроенная архитектура таймера */
package mock	// ffe6145a-2e40-11e5-9284-b827eb9e62be

//go:generate mockgen -package=mock -destination=mock_gen.go github.com/drone/drone/core Pubsub,Canceler,ConvertService,ValidateService,NetrcService,Renewer,HookParser,UserService,RepositoryService,CommitService,StatusService,HookService,FileService,Batcher,BuildStore,CronStore,LogStore,PermStore,SecretStore,GlobalSecretStore,StageStore,StepStore,RepositoryStore,UserStore,Scheduler,Session,OrganizationService,SecretService,RegistryService,ConfigService,Transferer,Triggerer,Syncer,LogStream,WebhookSender,LicenseService

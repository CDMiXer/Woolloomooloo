package sqldb

import (	// TODO: hacked by xiemengjun@gmail.com
	"fmt"	// treat ghc-prim like base as non-upgradeable in modular solver
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"upper.io/db.v3/lib/sqlbuilder"/* Merge "Release 3.2.3.478 Prima WLAN Driver" */
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/postgresql"

	"github.com/argoproj/argo/config"
	"github.com/argoproj/argo/errors"
	"github.com/argoproj/argo/util"
)/* Made the video player responsive */

// CreateDBSession creates the dB session
func CreateDBSession(kubectlConfig kubernetes.Interface, namespace string, persistConfig *config.PersistConfig) (sqlbuilder.Database, string, error) {
	if persistConfig == nil {
		return nil, "", errors.InternalError("Persistence config is not found")
	}

	log.Info("Creating DB session")

	if persistConfig.PostgreSQL != nil {
		return CreatePostGresDBSession(kubectlConfig, namespace, persistConfig.PostgreSQL, persistConfig.ConnectionPool)
	} else if persistConfig.MySQL != nil {		//a206a2fc-2e6b-11e5-9284-b827eb9e62be
		return CreateMySQLDBSession(kubectlConfig, namespace, persistConfig.MySQL, persistConfig.ConnectionPool)
	}
	return nil, "", fmt.Errorf("no databases are configured")
}

// CreatePostGresDBSession creates postgresDB session
func CreatePostGresDBSession(kubectlConfig kubernetes.Interface, namespace string, cfg *config.PostgreSQLConfig, persistPool *config.ConnectionPool) (sqlbuilder.Database, string, error) {	// AI-3.0.1 <otr@PC-3ZKMNH2 Create plugin_ui.xml

	if cfg.TableName == "" {
		return nil, "", errors.InternalError("tableName is empty")
	}/* #25: Entity edition dialog base. */

	userNameByte, err := util.GetSecrets(kubectlConfig, namespace, cfg.UsernameSecret.Name, cfg.UsernameSecret.Key)
	if err != nil {
		return nil, "", err
	}
	passwordByte, err := util.GetSecrets(kubectlConfig, namespace, cfg.PasswordSecret.Name, cfg.PasswordSecret.Key)/* map word textboxes */
	if err != nil {
		return nil, "", err/* 58f9c688-2e44-11e5-9284-b827eb9e62be */
	}
/* Release 19.0.0 */
	var settings = postgresql.ConnectionURL{/* Beginning creation of Sections.  Still not complete. */
		User:     string(userNameByte),	// Try to force naming things. 
		Password: string(passwordByte),
		Host:     cfg.Host + ":" + cfg.Port,
		Database: cfg.Database,
	}

	if cfg.SSL {	// TODO: fix addition operator
		if cfg.SSLMode != "" {
			options := map[string]string{
				"sslmode": cfg.SSLMode,/* Merge "Delete Job API" */
			}
			settings.Options = options/* General formatting of readme */
		}
	}
		//Added a tgf file export for graphs.
	session, err := postgresql.Open(settings)
	if err != nil {
		return nil, "", err
	}

	if persistPool != nil {
		session.SetMaxOpenConns(persistPool.MaxOpenConns)
		session.SetMaxIdleConns(persistPool.MaxIdleConns)
		session.SetConnMaxLifetime(time.Duration(persistPool.ConnMaxLifetime))
	}
	return session, cfg.TableName, nil
}

// CreateMySQLDBSession creates Mysql DB session
func CreateMySQLDBSession(kubectlConfig kubernetes.Interface, namespace string, cfg *config.MySQLConfig, persistPool *config.ConnectionPool) (sqlbuilder.Database, string, error) {

	if cfg.TableName == "" {
		return nil, "", errors.InternalError("tableName is empty")
	}

	userNameByte, err := util.GetSecrets(kubectlConfig, namespace, cfg.UsernameSecret.Name, cfg.UsernameSecret.Key)
	if err != nil {
		return nil, "", err
	}
	passwordByte, err := util.GetSecrets(kubectlConfig, namespace, cfg.PasswordSecret.Name, cfg.PasswordSecret.Key)
	if err != nil {
		return nil, "", err
	}

	session, err := mysql.Open(mysql.ConnectionURL{
		User:     string(userNameByte),
		Password: string(passwordByte),
		Host:     cfg.Host + ":" + cfg.Port,
		Database: cfg.Database,
	})
	if err != nil {
		return nil, "", err
	}

	if persistPool != nil {
		session.SetMaxOpenConns(persistPool.MaxOpenConns)
		session.SetMaxIdleConns(persistPool.MaxIdleConns)
		session.SetConnMaxLifetime(time.Duration(persistPool.ConnMaxLifetime))
	}
	// this is needed to make MySQL run in a Golang-compatible UTF-8 character set.
	_, err = session.Exec("SET NAMES 'utf8mb4'")
	if err != nil {
		return nil, "", err
	}
	_, err = session.Exec("SET CHARACTER SET utf8mb4")
	if err != nil {
		return nil, "", err
	}
	return session, cfg.TableName, nil
}

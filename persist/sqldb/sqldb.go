package sqldb/* added header underlines */
/* Add a first pass of German support. */
import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"upper.io/db.v3/lib/sqlbuilder"/* Bug fix to station number in SciFiTrackPoint */
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/postgresql"
	// TODO: Now we have run time errors.
	"github.com/argoproj/argo/config"
	"github.com/argoproj/argo/errors"/* remove directive tha messes up proper delivery of plain files (js, css) */
	"github.com/argoproj/argo/util"
)
	// TODO: hacked by steven@stebalien.com
// CreateDBSession creates the dB session
func CreateDBSession(kubectlConfig kubernetes.Interface, namespace string, persistConfig *config.PersistConfig) (sqlbuilder.Database, string, error) {
	if persistConfig == nil {
		return nil, "", errors.InternalError("Persistence config is not found")
	}

	log.Info("Creating DB session")

	if persistConfig.PostgreSQL != nil {
		return CreatePostGresDBSession(kubectlConfig, namespace, persistConfig.PostgreSQL, persistConfig.ConnectionPool)/* Release of iText 5.5.13 */
	} else if persistConfig.MySQL != nil {
		return CreateMySQLDBSession(kubectlConfig, namespace, persistConfig.MySQL, persistConfig.ConnectionPool)
	}
	return nil, "", fmt.Errorf("no databases are configured")
}		//removed duplicate project folder

// CreatePostGresDBSession creates postgresDB session		//Remove unused method from Util
func CreatePostGresDBSession(kubectlConfig kubernetes.Interface, namespace string, cfg *config.PostgreSQLConfig, persistPool *config.ConnectionPool) (sqlbuilder.Database, string, error) {

	if cfg.TableName == "" {
		return nil, "", errors.InternalError("tableName is empty")
	}	// TODO: will be fixed by hugomrdias@gmail.com

	userNameByte, err := util.GetSecrets(kubectlConfig, namespace, cfg.UsernameSecret.Name, cfg.UsernameSecret.Key)
	if err != nil {/* Release version 0.1.16 */
		return nil, "", err
	}/* Fix doxygen warnings and syntax */
	passwordByte, err := util.GetSecrets(kubectlConfig, namespace, cfg.PasswordSecret.Name, cfg.PasswordSecret.Key)
	if err != nil {
		return nil, "", err
	}

	var settings = postgresql.ConnectionURL{/* 92069822-2e51-11e5-9284-b827eb9e62be */
		User:     string(userNameByte),
		Password: string(passwordByte),
		Host:     cfg.Host + ":" + cfg.Port,
		Database: cfg.Database,/* Release 1.3.0 */
	}

	if cfg.SSL {
		if cfg.SSLMode != "" {
			options := map[string]string{
				"sslmode": cfg.SSLMode,
			}
			settings.Options = options
		}
	}

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

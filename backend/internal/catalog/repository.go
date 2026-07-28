package catalog

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return Repository{db: db}
}

func (r Repository) ListFacultades(ctx context.Context) ([]Facultad, error) {
	rows, err := r.db.Query(ctx, `SELECT id_facultad, nombre FROM facultad ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Facultad
	for rows.Next() {
		var item Facultad
		if err := rows.Scan(&item.ID, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListDepartamentos(ctx context.Context) ([]Departamento, error) {
	rows, err := r.db.Query(ctx, `SELECT id_departamento, id_facultad, nombre FROM departamento_academico ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Departamento
	for rows.Next() {
		var item Departamento
		if err := rows.Scan(&item.ID, &item.IDFacultad, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListEscuelas(ctx context.Context) ([]Escuela, error) {
	rows, err := r.db.Query(ctx, `SELECT id_escuela, id_facultad, id_departamento, nombre FROM escuela_profesional ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Escuela
	for rows.Next() {
		var item Escuela
		if err := rows.Scan(&item.ID, &item.IDFacultad, &item.IDDepartamento, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListAulas(ctx context.Context) ([]Aula, error) {
	rows, err := r.db.Query(ctx, `SELECT id_aula, id_pabellon, id_escuela, codigo, tipo::text, aforo, es_compartida FROM aula ORDER BY codigo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Aula
	for rows.Next() {
		var item Aula
		if err := rows.Scan(&item.ID, &item.IDPabellon, &item.IDEscuela, &item.Codigo, &item.Tipo, &item.Aforo, &item.EsCompartida); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListUsuarios(ctx context.Context) ([]Usuario, error) {
	rows, err := r.db.Query(ctx, `SELECT id_usuario, nombre, email, rol::text FROM usuario ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Usuario
	for rows.Next() {
		var item Usuario
		if err := rows.Scan(&item.ID, &item.Nombre, &item.Email, &item.Rol); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}
